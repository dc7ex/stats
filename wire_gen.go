// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/whj1990/go-core/store"
	"github.com/whj1990/go-core/trace"
	"google.golang.org/grpc"
	"mine/mine-grrpc/internal/repo"
	"mine/mine-grrpc/internal/service"
)

// Injectors from wire.go:

func initServer() (*grpc.Server, error) {
	loggerInterface := trace.NewGormLogger()
	db, err := store.NewDB(loggerInterface)
	if err != nil {
		return nil, err
	}
	reviewProjectRepo := repo.NewReviewProjectRepo(db)
	reviewService := service.NewReviewService(reviewProjectRepo)
	handleServerServer := newAppMineImpl(reviewService)
	server := newServer(handleServerServer)
	return server, nil
}