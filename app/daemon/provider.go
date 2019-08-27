package daemon

import (
	"context"
	"github.com/Nerufa/go-blueprint/app/db/domain"
	"github.com/Nerufa/go-blueprint/app/db/repo"
	"github.com/Nerufa/go-blueprint/app/db/trx"
	"github.com/Nerufa/go-shared/config"
	"github.com/Nerufa/go-shared/postgres"
	"github.com/Nerufa/go-shared/provider"
	"github.com/google/wire"
)

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

// Repo
type Repo struct {
	List domain.ListRepo
}

// Provider
func Provider(ctx context.Context, set provider.AwareSet, appSet AppSet, cfg Config) (*Daemon, func(), error) {
	g := New(ctx, set, appSet, cfg)
	return g, func() {}, nil
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
