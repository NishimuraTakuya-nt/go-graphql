package mapper

import (
	"github.com/NishimuraTakuya-nt/go-graphql/internal/adapters/primary/graphql/dto"
	"github.com/NishimuraTakuya-nt/go-graphql/internal/core/domain/models"
)

func ToDTO(domain *models.Sample) *dto.Sample {
	if domain == nil {
		return nil
	}
	return &dto.Sample{
		ID:        domain.ID,
		StringVal: domain.StringVal,
		IntVal:    domain.IntVal,
		ArrayVal:  domain.ArrayVal,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
