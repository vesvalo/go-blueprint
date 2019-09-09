package resolver

import (
	"context"
	"github.com/Nerufa/go-blueprint/generated/graphql"
	"github.com/Nerufa/go-blueprint/pkg/db/repo"
	"github.com/Nerufa/go-blueprint/pkg/db/trx"
	"github.com/Nerufa/go-shared/config"
	"github.com/Nerufa/go-shared/postgres"
	"github.com/Nerufa/go-shared/provider"
	"github.com/google/wire"
)

// Cfg
func Cfg(cfg config.Configurator) (*Config, func(), error) {
	c := &Config{}
	e := cfg.UnmarshalKeyOnReload(UnmarshalKey, c)
	return c, func() {}, e
}

// CfgTest
func CfgTest() (*Config, func(), error) {
	return &Config{}, func() {}, nil
}

type AppSet struct {
	Repo Repo
	Trx  *trx.Manager
}

// Provider
func Provider(ctx context.Context, set provider.AwareSet, appSet AppSet, cfg *Config) (graphql.Config, func(), error) {
	c := New(ctx, set, appSet, cfg)
	return c, func() {}, nil
}

var (
	ProviderRepo = wire.NewSet(
		repo.NewListRepo,
		trx.NewTrxManager,
	)
	ProviderRepoProduction = wire.NewSet(
		ProviderRepo,
		wire.Struct(new(Repo), "*"),
		postgres.WireSet,
	)
	ProviderTestRepo = wire.NewSet(
		ProviderRepo,
		wire.Struct(new(Repo), "*"),
		postgres.WireTestSet,
	)
	WireSet = wire.NewSet(
		Provider,
		Cfg,
		ProviderRepoProduction,
		wire.Struct(new(AppSet), "*"),
	)
	WireTestSet = wire.NewSet(
		Provider,
		CfgTest,
		ProviderTestRepo,
		wire.Struct(new(AppSet), "*"),
	)
)
