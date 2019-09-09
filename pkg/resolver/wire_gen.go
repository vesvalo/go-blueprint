// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package resolver

import (
	"context"
	"github.com/Nerufa/go-blueprint/generated/graphql"
	"github.com/Nerufa/go-blueprint/pkg/db/repo"
	"github.com/Nerufa/go-blueprint/pkg/db/trx"
	"github.com/Nerufa/go-shared/config"
	"github.com/Nerufa/go-shared/invoker"
	"github.com/Nerufa/go-shared/logger"
	"github.com/Nerufa/go-shared/metric"
	"github.com/Nerufa/go-shared/postgres"
	"github.com/Nerufa/go-shared/provider"
	"github.com/Nerufa/go-shared/tracing"
)

// Injectors from injector.go:

func Build(ctx context.Context, initial config.Initial, observer invoker.Observer) (graphql.Config, func(), error) {
	configurator, cleanup, err := config.Provider(initial, observer)
	if err != nil {
		return graphql.Config{}, nil, err
	}
	loggerConfig, cleanup2, err := logger.ProviderCfg(configurator)
	if err != nil {
		cleanup()
		return graphql.Config{}, nil, err
	}
	zap, cleanup3, err := logger.Provider(ctx, loggerConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	metricConfig, cleanup4, err := metric.ProviderCfg(configurator)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	scope, cleanup5, err := metric.Provider(ctx, zap, metricConfig)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	configuration, cleanup6, err := tracing.ProviderCfg(configurator)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	tracer, cleanup7, err := tracing.Provider(ctx, configuration, zap)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	awareSet := provider.AwareSet{
		Logger: zap,
		Metric: scope,
		Tracer: tracer,
	}
	postgresConfig, cleanup8, err := postgres.ProviderCfg(configurator)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	db, cleanup9, err := postgres.ProviderGORM(ctx, zap, postgresConfig)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	listRepo := repo.NewListRepo(db)
	resolverRepo := Repo{
		List: listRepo,
	}
	manager := trx.NewTrxManager(db)
	appSet := AppSet{
		Repo: resolverRepo,
		Trx:  manager,
	}
	resolverConfig, cleanup10, err := Cfg(configurator)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	graphqlConfig, cleanup11, err := Provider(ctx, awareSet, appSet, resolverConfig)
	if err != nil {
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	return graphqlConfig, func() {
		cleanup11()
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

func BuildTest(ctx context.Context, initial config.Initial, observer invoker.Observer) (graphql.Config, func(), error) {
	configurator, cleanup, err := config.Provider(initial, observer)
	if err != nil {
		return graphql.Config{}, nil, err
	}
	loggerConfig, cleanup2, err := logger.ProviderCfg(configurator)
	if err != nil {
		cleanup()
		return graphql.Config{}, nil, err
	}
	zap, cleanup3, err := logger.Provider(ctx, loggerConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	metricConfig, cleanup4, err := metric.ProviderCfg(configurator)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	scope, cleanup5, err := metric.Provider(ctx, zap, metricConfig)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	configuration, cleanup6, err := tracing.ProviderCfg(configurator)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	tracer, cleanup7, err := tracing.Provider(ctx, configuration, zap)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	awareSet := provider.AwareSet{
		Logger: zap,
		Metric: scope,
		Tracer: tracer,
	}
	db, cleanup8, err := postgres.ProviderGORMTest()
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	listRepo := repo.NewListRepo(db)
	resolverRepo := Repo{
		List: listRepo,
	}
	manager := trx.NewTrxManager(db)
	appSet := AppSet{
		Repo: resolverRepo,
		Trx:  manager,
	}
	resolverConfig, cleanup9, err := CfgTest()
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	graphqlConfig, cleanup10, err := Provider(ctx, awareSet, appSet, resolverConfig)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return graphql.Config{}, nil, err
	}
	return graphqlConfig, func() {
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
