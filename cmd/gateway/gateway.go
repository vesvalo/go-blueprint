package gateway

import (
	"context"
	"github.com/Nerufa/go-blueprint/cmd"
	"github.com/Nerufa/go-blueprint/pkg/http"
	"github.com/Nerufa/go-shared/config"
	"github.com/Nerufa/go-shared/logger"
	"github.com/spf13/cobra"
)

const Prefix = "cmd.gateway"

var (
	Cmd = &cobra.Command{
		Use:           "gateway",
		Short:         "Gateway API daemon",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(_ *cobra.Command, _ []string) {
			log := cmd.Slave.Logger().WithFields(logger.Fields{"service": Prefix})
			var (
				s *http.HTTP
				c func()
				e error
			)
			cmd.Slave.Executor(func(ctx context.Context, initial config.Initial) error {
				s, c, e = http.Build(ctx, initial)
				if e != nil {
					return e
				}
				c()
				return nil
			}, func(ctx context.Context) error {
				if e := s.ListenAndServe(); e != nil {
					return e
				}
				log.Info("daemon stopped successfully")
				return nil
			})
		},
	}
)

func init() {
	// pflags
	Cmd.PersistentFlags().StringP(http.UnmarshalKey+".bind", "b", ":8080", "bind address")
}
