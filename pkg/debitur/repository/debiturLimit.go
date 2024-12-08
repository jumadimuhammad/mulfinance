package repository

import (
	"context"
	"mulfinance/pkg/debitur/repository/entities"
)

func (u *debiturRepository) GetDebiturLimit(ctx context.Context, debiturID uint) (debiturLimits []entities.DebiturLimit, err error) {
	where := u.db.Where("debitur_id = ?", debiturID)
	if err = where.Find(&debiturLimits).Error; err != nil {
		return
	}

	return
}

func (u *debiturRepository) GetDebiturLimitByID(ctx context.Context, debiturLimitID uint) (debiturLimit *entities.DebiturLimit, err error) {
	if err = u.db.First(&debiturLimit, debiturLimitID).Error; err != nil {
		return
	}

	return
}
