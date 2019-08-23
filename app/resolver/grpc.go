package resolver

import (
	diGRPC "github.com/Nerufa/blueprint/app/grpc"
	"github.com/Nerufa/blueprint/cmd"
	"google.golang.org/grpc"
	"sync"
)

var (
	poolManager *diGRPC.PoolManager
	pmMu, svMu  sync.Mutex
	poolSrv     = map[string]*diGRPC.Pool{}
)

func getConnGRPC(srv string) (*grpc.ClientConn, diGRPC.Done, error) {
	if poolManager == nil {
		pmMu.Lock()
		defer pmMu.Unlock()
		if poolManager == nil {
			pm, _, e := diGRPC.Build(cmd.Slave)
			if e != nil {
				return nil, func() {}, e
			} else {
				poolManager = pm
			}
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
