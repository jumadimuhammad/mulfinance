package repository

import (
	"context"
	"mulfinance/pkg/limit/domain"
	"mulfinance/pkg/limit/ports"
	"mulfinance/pkg/limit/repository/entities"

	"gorm.io/gorm"
)

type limitRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) ports.IRepository {
	return &limitRepository{db}
}

func (u *limitRepository) ListLimit(ctx context.Context) (limits []entities.Limit, err error) {
	if err = u.db.Find(&limits).Error; err != nil {
		return
	}

	return
}

func (u *limitRepository) GetLimitByIDs(ctx context.Context, limidIDs []uint) (limits []entities.Limit, err error) {
	if err = u.db.Find(&limits, limidIDs).Error; err != nil {
		return
	}

	return
}

func mapToDomain(entities []entities.Limit) (limits []domain.Limit) {
	for _, entitie := range entities {
		limit := domain.Limit{
			ID:        entitie.ID,
			Limit:     entitie.Limit,
			Tenor:     entitie.Tenor,
			CreatedBy: entitie.CreatedBy,
			UpdatedBy: entitie.UpdatedBy,
			CreatedAt: entitie.CreatedAt,
			UpdatedAt: entitie.UpdatedAt,
		}
		limits = append(limits, limit)
	}

	return
}
