package usecases

import (
	"context"
	"time"

	"github.com/NishimuraTakuya-nt/go-graphql/internal/core/domain/models"
)

type SampleUsecase interface {
	Get(ctx context.Context, id string) (*models.Sample, error)
}

type sampleUsecase struct {
}

func NewSampleUsecase() SampleUsecase {
	return &sampleUsecase{}
}

func (u *sampleUsecase) Get(ctx context.Context, id string) (*models.Sample, error) {
	return &models.Sample{
		ID:        id,
		StringVal: "sample1",
		IntVal:    10,
		ArrayVal:  []string{"sample1", "sample2"},
		Email:     "aa@aa.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
