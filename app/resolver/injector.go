// +build wireinject

package resolver

import (
	"github.com/Nerufa/blueprint/generated/graphql"
	"github.com/Nerufa/go-shared/entrypoint"
	"github.com/Nerufa/go-shared/provider"
	"github.com/google/wire"
)

// Build
func Build(slave entrypoint.Slaver) (graphql.Config, func(), error) {
	panic(wire.Build(provider.Set, WireSet, wire.Struct(new(provider.AwareSet), "*")))
}

// Build
func BuildTest(slave entrypoint.Slaver) (graphql.Config, func(), error) {
	panic(wire.Build(provider.Set, WireTestSet, wire.Struct(new(provider.AwareSet), "*")))
}
