package http

import (
	"context"
	"github.com/Nerufa/blueprint/app/graphql"
	"github.com/Nerufa/go-shared/config"
	"github.com/Nerufa/go-shared/provider"
	"github.com/go-chi/chi"
	"github.com/google/wire"
	"github.com/rs/cors"
	"github.com/thoas/go-funk"
	"net/http"
	"net/url"
)

var mux *chi.Mux

// Cfg
func Cfg(cfg config.Configurator) (Config, func(), error) {
	c := Config{
		Debug: cfg.IsDebug(),
	}
	e := cfg.UnmarshalKey(unmarshalKey, &c)
	return c, func() {}, e
}

// CfgTest
func CfgTest() (Config, func(), error) {
	return Config{}, func() {}, nil
}

// Mux
func Mux(routers Routers, cfg Config) (*chi.Mux, func(), error) {
	if mux != nil {
		return mux, func() {}, nil
	}
	mux = chi.NewRouter()
	if cfg.Debug {
		mux.Use(cors.AllowAll().Handler)
	} else {
		m := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: false,
			AllowOriginRequestFunc: func(r *http.Request, origin string) bool {
				u, e := url.Parse(r.Header.Get("Origin"))
				if e != nil {
					return false
				}
				return funk.ContainsString(cfg.Cors.Allowed, u.Host)
			},
		})
		mux.Use(m.Handler)
	}
	// Middleware
	routers.GraphQL.Use(mux)
	// Routers
	routers.GraphQL.Routers(mux)
	return mux, func() {}, nil
}

// Routers
type Routers struct {
	GraphQL *graphql.GraphQL
}

var ProviderRouters = wire.NewSet(
	wire.Struct(new(Routers), "*"),
)

var ProviderRoutersTest = wire.NewSet(
	wire.Struct(new(Routers), "*"),
)

// Provider
func Provider(ctx context.Context, mux *chi.Mux, set provider.AwareSet, cfg Config) (*Http, func(), error) {
	g := New(ctx, mux, set, cfg)
	return g, func() {}, nil
}

var (
	WireSet     = wire.NewSet(Provider, Cfg, Mux, ProviderRouters, graphql.WireSet)
	WireTestSet = wire.NewSet(Provider, CfgTest, Mux, ProviderRoutersTest, graphql.WireTestSet)
)
