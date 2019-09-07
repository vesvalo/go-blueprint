package main

import (
	"github.com/Nerufa/go-blueprint/cmd/daemon"
	"github.com/Nerufa/go-blueprint/cmd/gateway"
	"github.com/Nerufa/go-blueprint/cmd/migrate"
	"github.com/Nerufa/go-blueprint/cmd/root"
	"github.com/Nerufa/go-blueprint/cmd/version"
)

func main() {
	root.Execute(gateway.Cmd, version.Cmd, migrate.Cmd, daemon.Cmd)
}
