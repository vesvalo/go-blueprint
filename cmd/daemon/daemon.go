package daemon

import (
	"context"
	"github.com/Nerufa/go-blueprint/cmd"
	"github.com/Nerufa/go-blueprint/pkg/daemon"
	"github.com/Nerufa/go-shared/entrypoint"
	"github.com/Nerufa/go-shared/logger"
	"github.com/spf13/cobra"
)

const Prefix = "cmd.deamon"

var (
	Cmd = &cobra.Command{
		Use:           "daemon",
		Short:         "GRPC API daemon",
		SilenceUsage:  true,
		SilenceErrors: true,
		Run: func(_ *cobra.Command, _ []string) {
			log := cmd.Slave.Logger().WithFields(logger.Fields{"service": Prefix})
			var (
				s *daemon.Daemon
				c func()
				e error
			)
			cmd.Slave.Executor(func(ctx context.Context) error {
				cmd.Slave.OnReload(func(ctx context.Context) {
					initial, ok := entrypoint.CtxExtractInitial(ctx)
					log.Info("catch reload in %s, debug: %v, ok: %v", logger.Args(Prefix, initial.WorkDir, ok))
				})
				initial, _ := entrypoint.CtxExtractInitial(ctx)
				s, c, e = daemon.Build(ctx, initial, cmd.Observer)
				if e != nil {
					return e
				}
				c()
				return nil
			}, func(ctx context.Context) error {
				if e := s.ListenAndServe(); e != nil {
					return e
				}
				return nil
			})
		},
	}
)

func init() {
	// pflags
	Cmd.PersistentFlags().StringP(daemon.UnmarshalKeyBind, "b", ":8081", "bind address")
}
