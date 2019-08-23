// +build wireinject

package grpc

import (
	"github.com/google/wire"
	"github.com/Nerufa/go-shared/entrypoint"
	"github.com/Nerufa/go-shared/provider"
)

// Build
func Build(slave entrypoint.Slaver) (*PoolManager, func(), error) {
	panic(wire.Build(provider.Set, WireSet, wire.Struct(new(provider.AwareSet), "*")))
}
