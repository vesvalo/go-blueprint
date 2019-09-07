package resolver

import (
	"context"
	"github.com/Nerufa/go-blueprint/cmd"
	diGRPC "github.com/Nerufa/go-blueprint/pkg/grpc"
	"google.golang.org/grpc"
	"sync"
)

var (
	poolManager *diGRPC.PoolManager
	pmMu, svMu  sync.Mutex
	poolSrv     = map[string]*diGRPC.Pool{}
)

func getConnGRPC(ctx context.Context, srv string) (*grpc.ClientConn, diGRPC.Done, error) {
	if poolManager == nil {
		pmMu.Lock()
		defer pmMu.Unlock()
		if poolManager == nil {
			pm, _, e := diGRPC.Build(ctx, cmd.Slave.Config())
			if e != nil {
				return nil, func() {}, e
			}
			poolManager = pm

		}
	}
	if _, ok := poolSrv[srv]; !ok {
		svMu.Lock()
		defer svMu.Unlock()
		if _, ok := poolSrv[srv]; !ok {
			p, _, e := poolManager.New(srv)
			if e != nil {
				return nil, func() {}, e
			}
			poolSrv[srv] = p
		}
	}
	return poolSrv[srv].Get()
}
