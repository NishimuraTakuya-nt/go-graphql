//go:build wireinject
// +build wireinject

package main

import (
	"github.com/NishimuraTakuya-nt/go-graphql/internal/adapters/primary/graphql/resolvers"
	"github.com/NishimuraTakuya-nt/go-graphql/internal/core/usecases"
	"github.com/google/wire"
)

func InitializeResolver() *resolvers.Resolver {
	wire.Build(
		usecases.Set,
		resolvers.Set,
	)
	return nil
}
