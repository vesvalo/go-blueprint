// +build wireinject

package http

import (
	"context"
	"github.com/Nerufa/go-shared/config"
	"github.com/Nerufa/go-shared/provider"
	"github.com/google/wire"
)

// Build
func Build(ctx context.Context, initial config.Initial) (*HTTP, func(), error) {
	panic(wire.Build(WireSet, provider.Set, wire.Struct(new(provider.AwareSet), "*")))
}

// BuildTest
func BuildTest(ctx context.Context, initial config.Initial) (*HTTP, func(), error) {
	panic(wire.Build(WireTestSet, provider.Set, wire.Struct(new(provider.AwareSet), "*")))
}
