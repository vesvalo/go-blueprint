// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package graphql

import (
	"context"
	"github.com/Nerufa/go-blueprint/pkg/db/repo"
	"github.com/Nerufa/go-blueprint/pkg/db/trx"
	"github.com/Nerufa/go-blueprint/pkg/resolver"
	"github.com/Nerufa/go-shared/config"
	"github.com/Nerufa/go-shared/logger"
	"github.com/Nerufa/go-shared/metric"
	"github.com/Nerufa/go-shared/postgres"
	"github.com/Nerufa/go-shared/provider"
	"github.com/Nerufa/go-shared/tracing"
)

// Injectors from injector.go:

func Build(ctx context.Context, initial config.Initial) (*GraphQL, func(), error) {
	configurator, cleanup, err := config.Provider(initial)
	if err != nil {
		return nil, nil, err
	}
	loggerConfig, cleanup2, err := logger.ProviderCfg(configurator)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	zap, cleanup3, err := logger.Provider(ctx, loggerConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metricConfig, cleanup4, err := metric.ProviderCfg(configurator)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	scope, cleanup5, err := metric.Provider(ctx, zap, metricConfig)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	configuration, cleanup6, err := tracing.ProviderCfg(configurator)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracer, cleanup7, err := tracing.Provider(ctx, configuration, zap)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
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
		return nil, nil, err
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
		return nil, nil, err
	}
	listRepo := repo.NewListRepo(db)
	resolverRepo := resolver.Repo{
		List: listRepo,
	}
	manager := trx.NewTrxManager(db)
	appSet := resolver.AppSet{
		Repo: resolverRepo,
		Trx:  manager,
	}
	resolverConfig, cleanup10, err := resolver.Cfg(configurator)
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
		return nil, nil, err
	}
	graphqlConfig, cleanup11, err := resolver.Provider(ctx, awareSet, appSet, resolverConfig)
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
		return nil, nil, err
	}
	config2, cleanup12, err := Cfg(configurator)
	if err != nil {
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
		return nil, nil, err
	}
	graphQL, cleanup13, err := Provider(ctx, graphqlConfig, awareSet, config2)
	if err != nil {
		cleanup12()
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
		return nil, nil, err
	}
	return graphQL, func() {
		cleanup13()
		cleanup12()
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

func BuildTest(ctx context.Context, initial config.Initial) (*GraphQL, func(), error) {
	configurator, cleanup, err := config.Provider(initial)
	if err != nil {
		return nil, nil, err
	}
	loggerConfig, cleanup2, err := logger.ProviderCfg(configurator)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	zap, cleanup3, err := logger.Provider(ctx, loggerConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metricConfig, cleanup4, err := metric.ProviderCfg(configurator)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	scope, cleanup5, err := metric.Provider(ctx, zap, metricConfig)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	configuration, cleanup6, err := tracing.ProviderCfg(configurator)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracer, cleanup7, err := tracing.Provider(ctx, configuration, zap)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
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
		return nil, nil, err
	}
	listRepo := repo.NewListRepo(db)
	resolverRepo := resolver.Repo{
		List: listRepo,
	}
	manager := trx.NewTrxManager(db)
	appSet := resolver.AppSet{
		Repo: resolverRepo,
		Trx:  manager,
	}
	resolverConfig, cleanup9, err := resolver.CfgTest()
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	graphqlConfig, cleanup10, err := resolver.Provider(ctx, awareSet, appSet, resolverConfig)
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
		return nil, nil, err
	}
	config2, cleanup11, err := CfgTest()
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
		return nil, nil, err
	}
	graphQL, cleanup12, err := Provider(ctx, graphqlConfig, awareSet, config2)
	if err != nil {
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
		return nil, nil, err
	}
	return graphQL, func() {
		cleanup12()
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