package ports

import (
	"context"
	"mulfinance/pkg/debitur/repository/entities"
)

type IRepository interface {
	ListDebitur(ctx context.Context) (debiturs []entities.Debitur, err error)
	GetDebiturByID(ctx context.Context, debiturID uint) (debitur *entities.Debitur, err error)
	GetDebiturLimit(ctx context.Context, debiturID uint) (debiturLimit []entities.DebiturLimit, err error)
	GetDebiturLimitByID(ctx context.Context, debiturLimitID uint) (debiturLimit *entities.DebiturLimit, err error)
}
