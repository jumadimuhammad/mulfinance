package ports

import (
	"context"
	"mulfinance/pkg/transaction/domain"
)

type IUsecase interface {
	CreateTransaction(ctx context.Context, transaction domain.Transaction) (err error)
	ListTransaction(ctx context.Context) (transactions []domain.TransactionResponse, err error)
}
