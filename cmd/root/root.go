package root

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Nerufa/go-blueprint/cmd"
	"github.com/Nerufa/go-blueprint/cmd/version"
	"github.com/Nerufa/go-shared/config"
	"github.com/Nerufa/go-shared/entrypoint"
	"github.com/Nerufa/go-shared/logger"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

var (
	initial = config.Initial{}
	log     logger.Logger
	ep      entrypoint.Master
	e       error
	c       func()
)

const (
	prefix               = "cmd.root"
	envPrefix            = "APP"
	envWorkDir           = "APP_WD"
	viperCfgType         = "yaml"
	defaultGracefulDelay = 50 * time.Millisecond
)

// http://www.patorjk.com/software/taag/#p=display&f=Big&t=Blueprint

var logo = `
  ____  _                       _       _   
 |  _ \| |                     (_)     | |  
 | |_) | |_   _  ___ _ __  _ __ _ _ __ | |_ 
 |  _ <| | | | |/ _ \ '_ \| '__| | '_ \| __|
 | |_) | | |_| |  __/ |_) | |  | | | | | |_ 
 |____/|_|\__,_|\___| .__/|_|  |_|_| |_|\__|
                    | |                     
                    |_|      

		VERSION: %v`

// Root command
var rootCmd = &cobra.Command{
	Use:           "bin [command]",
	Long:          "",
	Short:         fmt.Sprintf(logo, version.Version()),
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRunE: func(subCmd *cobra.Command, _ []string) error {

		// initializing
		initial.WorkDir = os.Getenv(envWorkDir)
		if len(initial.WorkDir) == 0 {
			initial.WorkDir, e = filepath.Abs(filepath.Dir(os.Args[0]))
			if e != nil {
				return e
			}
		}
		initial.WorkDir, e = filepath.Abs(initial.WorkDir)
		if e != nil {
			return e
		}

		// bin pflags to viper
		e = initial.Viper.BindPFlags(subCmd.Parent().PersistentFlags())
		if e != nil {
			return e
		}
		e = initial.Viper.BindPFlags(subCmd.PersistentFlags())
		if e != nil {
			return e
		}

		initial.Viper.SetConfigFile(initial.ConfigPath)

		if initial.ConfigPath != "" {
			e := initial.Viper.ReadInConfig()
			if e != nil {
				return fmt.Errorf("can't read config, %v", errors.WithMessage(e, prefix))
			}
		}

		ep, c, e = entrypoint.Build(context.Background(), initial)
		if e != nil {
			return e
		}
		defer c()

		cmd.Slave = ep
		log = ep.Logger().WithFields(logger.Fields{"service": prefix})

		go func() {
			reloadSignal := make(chan os.Signal, 1)
			signal.Notify(reloadSignal, syscall.SIGHUP)
			for {
				sig := <-reloadSignal
				ep.Reload()
				ep.Logger().Info("OS signaled `%v`, reload", logger.Args(sig.String()))
			}
		}()

		go func() {
			shutdownSignal := make(chan os.Signal, 1)
			signal.Notify(shutdownSignal, syscall.SIGTERM, syscall.SIGINT)
			sig := <-shutdownSignal
			ep.Logger().Info("OS signaled `%v`, graceful shutdown in %s", logger.Args(sig.String(), initial.GracefulDelay), logger.WithTags(logger.Tags{"test"}))
			ctx, _ := context.WithTimeout(context.Background(), initial.GracefulDelay)
			ep.Shutdown(ctx, 0)
		}()

		return nil
	},
	PersistentPostRun: func(_ *cobra.Command, _ []string) {

		preRun := func() error {
			if initial.Debug {
				fmt.Printf(logo, version.Version())
				fmt.Println(color.RedString("\n\n# DEBUG INFO\n"))
				fmt.Printf("\nWork directory: %v\n\n", ep.WorkDir())
				fmt.Println(color.GreenString("# CONFIG FILE SETTINGS\n\n"))
				b, _ := json.Marshal(initial.Viper.AllSettings())
				var out bytes.Buffer
				e := json.Indent(&out, b, "", "  ")
				if e != nil {
					log.Error("can't prettify config")
					os.Exit(1)
				}
				fmt.Println(out.String())
				fmt.Println(color.CyanString("\n# LOGS\n\n"))
			}
			_, err := maxprocs.Set(maxprocs.Logger(log.Printf))
			return err
		}

		if e = ep.Serve(preRun); e != nil {
			_ = preRun()
			log.Error(e.Error())
			ep.Shutdown(context.Background(), 1)
		}
	},
}

func init() {
	initial.Viper = viper.New()
	initial.Viper.SetConfigType(viperCfgType)
	initial.Viper.SetEnvPrefix(envPrefix)
	initial.Viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	initial.Viper.AutomaticEnv()
	// pflags
	rootCmd.PersistentFlags().StringVarP(&initial.ConfigPath, "cmd.root.config", "c", "", "config file")
	rootCmd.PersistentFlags().BoolVarP(&initial.Debug, "cmd.root.debug.power", "d", false, "debug mode")
	rootCmd.PersistentFlags().StringVarP(&initial.LoggerLevel, "cmd.root.level", "l", "info", "logger level")
	rootCmd.PersistentFlags().StringSliceVarP(&initial.DebugTags, "cmd.root.debug.tags", "t", []string{}, "logger tags for filter output, e.g.: -t tag -t tag2 -t key:value")
	rootCmd.PersistentFlags().DurationVar(&initial.GracefulDelay, "cmd.root.graceful.delay", defaultGracefulDelay, "graceful delay")
}

func Execute(cmds ...*cobra.Command) {
	rootCmd.AddCommand(cmds...)
	if e := rootCmd.Execute(); e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", e.Error())
		os.Exit(1)
	}
}
