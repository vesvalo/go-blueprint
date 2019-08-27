package daemon

import (
	"github.com/Nerufa/go-blueprint/app/daemon"
	"github.com/Nerufa/go-blueprint/cmd"
	"github.com/Nerufa/go-shared/logger"
	"github.com/spf13/cobra"
)

var (
	bind string
	Cmd  = &cobra.Command{
		Use:           "daemon",
		Short:         "GRPC API daemon",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(_ *cobra.Command, _ []string) {
			var (
				e error
				s *daemon.Daemon
				c func()
			)
			log, c, e := logger.Build(cmd.Slave)
			if e != nil {
				panic(e)
			}
			defer c()
			defer func() {
				if r := recover(); r != nil {
					if re, _ := r.(error); re != nil {
						log.Error(re.Error())
					} else {
						log.Alert("unhandled panic, err: %v", logger.Args(r))
					}
				}
			}()
			s, c, e = daemon.Build(cmd.Slave)
			if e != nil {
				log.Error(e.Error())
				return
			}
			defer c()
			if err := s.ListenAndServe(bind); err != nil {
				log.Error(err.Error())
				return
			}
			log.Info("daemon stopped successfully")
		},
	}
)

func init() {
	// pflags
	Cmd.PersistentFlags().StringVarP(&bind, "bind", "b", ":8081", "bind address")
}
