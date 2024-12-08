package repository

import (
	"context"
	"mulfinance/pkg/transaction/ports"
	"mulfinance/pkg/transaction/repository/entities"

	"gorm.io/gorm"
)

type transactionRepository struct{}

func NewRepository() ports.IRepository {
	return &transactionRepository{}
}

func (u *transactionRepository) CreateTransaction(ctx context.Context, db *gorm.DB, transaction entities.Transaction) (err error) {
	if err = db.Create(&transaction).Error; err != nil {
		return
	}

	return
}

func (u *transactionRepository) ListTransaction(ctx context.Context, db *gorm.DB) (transactions []entities.Transaction, err error) {
	if err = db.Find(&transactions).Error; err != nil {
		return
	}

	return
}

func (u *transactionRepository) GetTransactionByDebiturID(ctx context.Context, db *gorm.DB, debiturID uint) (transactions []entities.Transaction, err error) {
	result := db.Joins("INNER JOIN transaction_limits ON transactions.transaction_limit_id = transaction_limits.id").
		Where("debitur_id = ?", debiturID).
		Find(&transactions)

	if err = result.Error; err != nil {
		return
	}

	return
}
