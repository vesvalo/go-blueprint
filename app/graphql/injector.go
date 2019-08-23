// +build wireinject

package graphql

import (
	"github.com/google/wire"
	"github.com/Nerufa/go-shared/entrypoint"
	"github.com/Nerufa/go-shared/provider"
)

// Build
func Build(slave entrypoint.Slaver) (*GraphQL, func(), error) {
	panic(wire.Build(provider.Set, WireSet, wire.Struct(new(provider.AwareSet), "*")))
}

// BuildTest
func BuildTest(slave entrypoint.Slaver) (*GraphQL, func(), error) {
	panic(wire.Build(provider.Set, WireTestSet, wire.Struct(new(provider.AwareSet), "*")))
}