// +build wireinject

package daemon

import (
	"github.com/Nerufa/go-shared/entrypoint"
	"github.com/Nerufa/go-shared/provider"
	"github.com/google/wire"
)

// Build
func Build(slave entrypoint.Slaver) (*Daemon, func(), error) {
	panic(wire.Build(WireSet, provider.Set, wire.Struct(new(provider.AwareSet), "*")))
}

// BuildTest
func BuildTest(slave entrypoint.Slaver) (*Daemon, func(), error) {
	panic(wire.Build(WireTestSet, provider.Set, wire.Struct(new(provider.AwareSet), "*")))
}
