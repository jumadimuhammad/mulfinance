package ports

import (
	"context"
	"mulfinance/pkg/limit/repository/entities"
)

type IRepository interface {
	ListLimit(ctx context.Context) (limits []entities.Limit, err error)
	GetLimitByIDs(ctx context.Context, limitIDs []uint) (limits []entities.Limit, err error)
}
