//go:build wireinject
// +build wireinject

package main

import (
	"github.com/dc7ex/stats/internal/repo"
	"github.com/dc7ex/stats/internal/service"

	"google.golang.org/grpc"

	"github.com/google/wire"
)

func initServer() (*grpc.Server, error) {
	panic(wire.Build(repo.ProviderSet, service.ProviderSet, newAppImpl, newServer))
}
