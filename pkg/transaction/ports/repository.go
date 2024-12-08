package ports

import (
	"context"
	"mulfinance/pkg/transaction/repository/entities"

	"gorm.io/gorm"
)

type IRepository interface {
	CreateTransaction(ctx context.Context, db *gorm.DB, transaction entities.Transaction) (err error)
	ListTransaction(ctx context.Context, db *gorm.DB) (transactions []entities.Transaction, err error)
	GetTransactionByDebiturID(ctx context.Context, db *gorm.DB, debiturID uint) (transactions []entities.Transaction, err error)

	CreateTransactionLimit(ctx context.Context, db *gorm.DB, transactionLimit entities.TransactionLimit) (transactionLimitID uint, err error)
	GetTransactionLimitProcessedByDebiturID(ctx context.Context, db *gorm.DB, debiturID uint) (transactions *entities.TransactionLimit, err error)
	GetTransactionLimitProcessed(ctx context.Context, db *gorm.DB, debiturID uint) (transactionLimitTotal *entities.TransactionLimitTotal, err error)
}
