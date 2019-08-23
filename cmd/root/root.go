package root

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Nerufa/blueprint/cmd"
	"github.com/Nerufa/blueprint/cmd/daemon"
	"github.com/Nerufa/blueprint/cmd/gateway"
	"go.uber.org/automaxprocs/maxprocs"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/Nerufa/blueprint/cmd/migrate"
	"github.com/Nerufa/blueprint/cmd/version"
	"github.com/Nerufa/go-shared/entrypoint"
	"github.com/Nerufa/go-shared/logger"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg        = entrypoint.Config{}
	configPath string
	v          *viper.Viper
	log        logger.Logger
	ep         entrypoint.Master
)

const prefix = "cmd.root"

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
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		l, c, e := logger.Build(ep)
		if e != nil {
			panic(e)
		}
		defer c()

		log = l.WithFields(logger.Fields{"service": prefix})
		v.SetConfigFile(configPath)

		if configPath != "" {
			e := v.ReadInConfig()
			if e != nil {
				log.Error("can't read config, %v", logger.Args(errors.WithMessage(e, prefix)))
				os.Exit(1)
			}
		}

		if cfg.Debug {
			fmt.Printf(logo, version.Version())
			fmt.Println(color.RedString("\n# DEBUG INFO\n"))
			fmt.Printf("\nWork directory: %v\n\n", ep.WorkDir())
			fmt.Println(color.GreenString("# CONFIG FILE SETTINGS\n\n"))
			b, _ := json.Marshal(v.AllSettings())
			var out bytes.Buffer
			e = json.Indent(&out, b, "", "  ")
			if e != nil {
				log.Error("can't prettify config")
				os.Exit(1)
			}
			fmt.Println(string(out.Bytes()))
			fmt.Println(color.CyanString("\n# LOGS\n\n"))
		}

		_, _ = maxprocs.Set(maxprocs.Logger(log.Printf))
	},
}

func init() {
	v = viper.New()
	v.SetConfigType("yaml")
	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	v.AutomaticEnv()

	// pflags
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file")
	rootCmd.PersistentFlags().BoolVarP(&cfg.Debug, "debug", "d", false, "debug mode")
	rootCmd.PersistentFlags().StringVarP(&cfg.LoggerLevel, "level", "l", "info", "logger level")
	rootCmd.PersistentFlags().StringSliceVarP(&cfg.DebugTags, "debug.tags", "t", []string{}, "logger tags for filter output, e.g.: -t tag -t tag2 -t key:value")
	rootCmd.PersistentFlags().DurationVar(&cfg.GracefulDelay, "graceful.delay", 50*time.Millisecond, "graceful delay")

	// initializing
	wd := os.Getenv("APP_WD")
	if len(wd) == 0 {
		wd, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}
	wd, _ = filepath.Abs(wd)
	cfg.Viper = v
	ep, _ = entrypoint.Initialize(&cfg)
	cmd.Slave = ep

	// bin pflags to viper
	_ = v.BindPFlags(rootCmd.PersistentFlags())

	go func() {
		reloadSignal := make(chan os.Signal)
		signal.Notify(reloadSignal, syscall.SIGHUP)
		for {
			sig := <-reloadSignal
			ep.Reload()
			log.Info("OS signaled `%v`, reload", logger.Args(sig.String()))
		}
	}()

	go func() {
		shutdownSignal := make(chan os.Signal)
		signal.Notify(shutdownSignal, syscall.SIGTERM, syscall.SIGINT)
		sig := <-shutdownSignal
		log.Info("OS signaled `%v`, graceful shutdown in %s", logger.Args(sig.String(), cfg.GracefulDelay), logger.WithTags(logger.Tags{"test"}))
		ctx, _ := context.WithTimeout(context.Background(), cfg.GracefulDelay)
		ep.Shutdown(ctx, 0)
	}()
}

func Execute() {
	rootCmd.AddCommand(gateway.Cmd, version.Cmd, migrate.Cmd, daemon.Cmd)
	if e := rootCmd.Execute(); e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", e.Error())
		os.Exit(1)
	}
}
