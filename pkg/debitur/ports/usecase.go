package ports

import (
	"context"
	"mulfinance/pkg/debitur/domain"
	transactionDomain "mulfinance/pkg/transaction/domain"
)

type IUsecase interface {
	ListDebitur(ctx context.Context) (debitur []domain.DebiturResponse, err error)
	GetDebiturByID(ctx context.Context, debiturID uint) (debitur *domain.DebiturResponse, err error)
	GetDebiturLimit(ctx context.Context, debiturID uint) (debitur domain.DebiturLimitResponse, err error)
	GetDebiturTransaction(ctx context.Context, debiturID uint) (transactions []transactionDomain.TransactionResponse, err error)
}
