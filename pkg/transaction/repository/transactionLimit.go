package repository

import (
	"context"
	"errors"
	"mulfinance/pkg/transaction/repository/entities"

	"gorm.io/gorm"
)

func (u *transactionRepository) CreateTransactionLimit(ctx context.Context, db *gorm.DB, transactionLimit entities.TransactionLimit) (transactionLimitID uint, err error) {
	if err = db.Create(&transactionLimit).Error; err != nil {
		return
	}

	transactionLimitID = transactionLimit.ID

	return
}

func (u *transactionRepository) GetTransactionLimitProcessedByDebiturID(ctx context.Context, db *gorm.DB, debiturID uint) (transactionLimit *entities.TransactionLimit, err error) {
	where := db.Where("debitur_id = ? AND status = ?", debiturID, "processed")
	err = where.First(&transactionLimit).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
			return
		}

		return
	}

	return
}

func (u *transactionRepository) GetTransactionLimitProcessed(ctx context.Context, db *gorm.DB, debiturID uint) (transactionLimitTotal *entities.TransactionLimitTotal, err error) {
	query := db.Table("transaction_limits AS tl").
		Select(`
		l.limit AS credit_limit, 
		SUM(COALESCE(t.otr, 0)) AS credit_total
	`).
		Joins(`
		INNER JOIN debitur_limits AS dl ON tl.debitur_limit_id = dl.id
		INNER JOIN limits AS l ON dl.limit_id = l.id
		LEFT JOIN transactions AS t ON tl.id = t.transaction_limit_id AND tl.status = ?
	`, "processed").
		Where("dl.debitur_id = ?", debiturID).
		Group("tl.id, l.limit").
		Scan(&transactionLimitTotal)

	if err = query.Error; err != nil {
		return
	}

	return
}
