package grpc

import (
	"context"
	"github.com/Nerufa/go-shared/provider"
	"github.com/Nerufa/go-shared/tracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// PoolManager
type PoolManager struct {
	ctx     context.Context
	tracing tracing.Tracer
	cfg     *Config
}

// New
func (p *PoolManager) New(service string) (_ *Pool, loaded bool, _ error) {
	s, ok := p.cfg.Services[service]
	if !ok {
		return nil, false, errCfgInvalid
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	cl := s.ClientParameters
	if cl == nil {
		if p.cfg.ClientParameters != nil {
			cl = p.cfg.ClientParameters
		} else {
			cl = &keepalive.ClientParameters{}
		}
	}
	opts = append(opts, grpc.WithKeepaliveParams(*cl))
	pool, l := NewPool(p.ctx, service, s.Target,
		MaxConn(s.MaxConn),
		InitConn(s.InitConn),
		MaxLifeDuration(s.MaxLifeDuration),
		IdleTimeout(s.IdleTimeout),
		ConnOptions(opts...),
	)
	return pool, l, nil
}

// NewPoolManager
func NewPoolManager(ctx context.Context, set provider.AwareSet, cfg *Config) *PoolManager {
	return &PoolManager{ctx: ctx, tracing: set.Tracer, cfg: cfg}
}
