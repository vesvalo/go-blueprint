package graphql

import (
	"context"
	"github.com/Nerufa/go-blueprint/generated/graphql"
	"github.com/Nerufa/go-blueprint/pkg/resolver"
	"github.com/Nerufa/go-shared/config"
	"github.com/Nerufa/go-shared/provider"
	"github.com/google/wire"
	"net/http"
)

// Cfg
func Cfg(cfg config.Configurator) (*Config, func(), error) {
	c := &Config{}
	e := cfg.UnmarshalKeyOnReload(UnmarshalKey, c)
	c.Middleware = []func(http.Handler) http.Handler{
	}
	return c, func() {}, e
}

// CfgTest
func CfgTest() (*Config, func(), error) {
	return &Config{}, func() {}, nil
}

// Provider
func Provider(ctx context.Context, resolver graphql.Config, set provider.AwareSet, cfg *Config) (*GraphQL, func(), error) {
	g := New(ctx, resolver, set, cfg)
	return g, func() {}, nil
}

var (
	WireSet     = wire.NewSet(Provider, Cfg, resolver.WireSet)
	WireTestSet = wire.NewSet(Provider, CfgTest, resolver.WireTestSet)
)
