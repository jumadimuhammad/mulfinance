package repository

import (
	"context"
	"mulfinance/pkg/debitur/domain"
	"mulfinance/pkg/debitur/ports"
	"mulfinance/pkg/debitur/repository/entities"

	"gorm.io/gorm"
)

type debiturRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) ports.IRepository {
	return &debiturRepository{db}
}

func (u *debiturRepository) ListDebitur(ctx context.Context) (debiturs []entities.Debitur, err error) {
	if err = u.db.Find(&debiturs).Error; err != nil {
		return
	}

	return
}

func (u *debiturRepository) GetDebiturByID(ctx context.Context, debiturID uint) (debitur *entities.Debitur, err error) {
	if err = u.db.First(&debitur, debiturID).Error; err != nil {
		return
	}

	return
}

func mapToDomain(entities []entities.Debitur) (debiturs []domain.Debitur) {
	for _, entitie := range entities {
		debitur := domain.Debitur{
			ID:          entitie.ID,
			NIK:         entitie.NIK,
			FullName:    entitie.FullName,
			LegalName:   entitie.LegalName,
			BirthPlace:  entitie.BirthPlace,
			BirthDate:   entitie.BirthDate,
			Salary:      entitie.Salary,
			KTPPhoto:    entitie.KTPPhoto,
			SelfiePhoto: entitie.SelfiePhoto,
			CreatedBy:   entitie.CreatedBy,
			UpdatedBy:   entitie.UpdatedBy,
			CreatedAt:   entitie.CreatedAt,
			UpdatedAt:   entitie.UpdatedAt,
		}
		debiturs = append(debiturs, debitur)
	}

	return
}
