package ports

import (
	"context"
	"mulfinance/pkg/limit/domain"
)

type IUsecase interface {
	ListLimit(ctx context.Context) (limits []domain.Limit, err error)
}
