package resolvers

import "github.com/NishimuraTakuya-nt/go-graphql/internal/core/usecases"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	sampleUsecase usecases.SampleUsecase
}

func NewResolver(sampleUsecase usecases.SampleUsecase) *Resolver {
	return &Resolver{
		sampleUsecase: sampleUsecase,
	}
}
