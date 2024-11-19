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

func ToDomain(dto *dto.Sample) *models.Sample {
	if dto == nil {
		return nil
	}
	return &models.Sample{
		ID:        dto.ID,
		StringVal: dto.StringVal,
		IntVal:    dto.IntVal,
		ArrayVal:  dto.ArrayVal,
		Email:     dto.Email,
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}

func ToDomainCreateSampleInput(dto *dto.CreateSampleInput) *models.CreateSampleInput {
	if dto == nil {
		return nil
	}
	return &models.CreateSampleInput{
		StringVal: dto.StringVal,
		IntVal:    dto.IntVal,
		ArrayVal:  dto.ArrayVal,
		Email:     dto.Email,
	}
}
