package usecases

import (
	"context"
	"time"

	"github.com/NishimuraTakuya-nt/go-graphql/internal/core/domain/models"
)

type SampleUsecase interface {
	Get(ctx context.Context, id string) (*models.Sample, error)
	List(ctx context.Context, offset, limit int) ([]*models.Sample, error)
	Create(ctx context.Context, sample *models.CreateSampleInput) (*models.Sample, error)
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

func (u *sampleUsecase) List(ctx context.Context, offset, limit int) ([]*models.Sample, error) {
	return []*models.Sample{
		{
			ID:        "1",
			StringVal: "sample1",
			IntVal:    10,
			ArrayVal:  []string{"sample1", "sample2"},
			Email:     "aa@aa.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "2",
			StringVal: "sample2",
			IntVal:    11,
			ArrayVal:  []string{"sample3", "sample4"},
			Email:     "bb@bb.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}

func (u *sampleUsecase) Create(ctx context.Context, input *models.CreateSampleInput) (*models.Sample, error) {
	return &models.Sample{
		ID:        "sample1",
		StringVal: input.StringVal,
		IntVal:    input.IntVal,
		ArrayVal:  input.ArrayVal,
		Email:     input.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
