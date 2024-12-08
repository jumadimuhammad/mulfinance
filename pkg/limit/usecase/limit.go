package usecase

import (
	"context"
	"mulfinance/pkg/limit/domain"
	"mulfinance/pkg/limit/ports"
)

type limitUsecase struct {
	limitRepository ports.IRepository
}

func NewUsecase(dRepo ports.IRepository) ports.IUsecase {
	return &limitUsecase{dRepo}
}

func (d *limitUsecase) ListLimit(ctx context.Context) (limits []domain.Limit, err error) {
	resp, err := d.limitRepository.ListLimit(ctx)
	if err != nil {
		return nil, err
	}

	for _, entitie := range resp {
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
