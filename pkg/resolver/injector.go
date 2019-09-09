// +build wireinject

package resolver

import (
	"context"
	"github.com/Nerufa/go-blueprint/generated/graphql"
	"github.com/Nerufa/go-shared/config"
	"github.com/Nerufa/go-shared/invoker"
	"github.com/Nerufa/go-shared/provider"
	"github.com/google/wire"
)

// Build
func Build(ctx context.Context, initial config.Initial, observer invoker.Observer) (graphql.Config, func(), error) {
	panic(wire.Build(provider.Set, WireSet, wire.Struct(new(provider.AwareSet), "*")))
}

// Build
func BuildTest(ctx context.Context, initial config.Initial, observer invoker.Observer) (graphql.Config, func(), error) {
	panic(wire.Build(provider.Set, WireTestSet, wire.Struct(new(provider.AwareSet), "*")))
}
