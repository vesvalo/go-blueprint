package grpc

import (
	"context"
	"github.com/Nerufa/go-shared/config"
	"github.com/Nerufa/go-shared/provider"
	"github.com/google/wire"
	"google.golang.org/grpc/keepalive"
	"sync"
	"time"
)

var (
	pm   *PoolManager
	muPM sync.Mutex
)

// Cfg
func Cfg(cfg config.Configurator) (*Config, func(), error) {
	c := &Config{}
	e := cfg.UnmarshalKey(UnmarshalKey, c)
	return c, func() {}, e
}

// Service
type Service struct {
	Target           string
	MaxConn          int
	InitConn         int
	MaxLifeDuration  time.Duration
	IdleTimeout      time.Duration
	ClientParameters *keepalive.ClientParameters
}

// Config
type Config struct {
	Services         map[string]*Service
	ClientParameters *keepalive.ClientParameters
}

// Provider
func Provider(ctx context.Context, set provider.AwareSet, cfg *Config) (*PoolManager, func(), error) {
	muPM.Lock()
	defer muPM.Unlock()
	if pm != nil {
		return pm, func() {}, nil
	}
	pm = NewPoolManager(ctx, set, cfg)
	return pm, func() {}, nil
}

var WireSet = wire.NewSet(Provider, Cfg)
