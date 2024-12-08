package usecase

import (
	"context"
	"errors"
	"fmt"
	debiturPorts "mulfinance/pkg/debitur/ports"
	"mulfinance/pkg/transaction/domain"
	"mulfinance/pkg/transaction/ports"
	"mulfinance/pkg/transaction/repository/entities"
	"time"

	"math/rand"

	"gorm.io/gorm"
)

type transactionUsecase struct {
	transactionRepository ports.IRepository
	debiturRepository     debiturPorts.IRepository
	connection            *gorm.DB
}

func NewUsecase(transRepo ports.IRepository, debRepo debiturPorts.IRepository, connection *gorm.DB) ports.IUsecase {
	return &transactionUsecase{transRepo, debRepo, connection}
}

func (d *transactionUsecase) CreateTransaction(ctx context.Context, data domain.Transaction) (err error) {
	tx := d.connection.Begin()

	respDebLimit, err := d.debiturRepository.GetDebiturLimit(ctx, data.DebiturID)
	if err != nil {
		return errors.New("debitur don't have limit")
	}

	if len(respDebLimit) == 0 {
		return errors.New("debitur don't have limit")
	}

	for _, limit := range respDebLimit {
		if limit.ID == data.DebiturLimitID {

			break
		}

		return errors.New("debitur don't have limit")
	}

	trans, err := d.transactionRepository.GetTransactionLimitProcessed(ctx, tx, data.DebiturID)
	if err != nil {
		return err
	}

	if trans != nil {
		if (trans.Limit - trans.TotalCredit) < data.OTR {
			limit := trans.Limit - trans.TotalCredit
			errN := fmt.Sprintf("limit for transaction, transaction under %v", limit)
			return errors.New(errN)
		}
	}

	resp, err := d.transactionRepository.GetTransactionLimitProcessedByDebiturID(ctx, tx, data.DebiturID)
	if err != nil {
		return err
	}

	transactionLimitID := resp.ID
	if resp.ID == 0 {
		entitieTransactionLimit := entities.TransactionLimit{
			DebiturID:      data.DebiturID,
			DebiturLimitID: data.DebiturLimitID,
			Status:         "processed",
			CreatedBy:      data.CreatedBy,
			UpdatedBy:      data.UpdatedBy,
		}

		transactionLimitID, err = d.transactionRepository.CreateTransactionLimit(ctx, tx, entitieTransactionLimit)
		if err != nil {
			return err
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(100000000)
	contractNumber := fmt.Sprintf("00%08d", randomNumber)

	entitieTransaction := entities.Transaction{
		TransactionLimitID: transactionLimitID,
		ContractNumber:     contractNumber,
		OTR:                data.OTR,
		AdminFee:           data.AdminFee,
		InstallmentAmount:  data.InstallmentAmount,
		AmountInterest:     data.AmountInterest,
		ProductName:        data.ProductName,
		CreatedBy:          data.CreatedBy,
		UpdatedBy:          data.UpdatedBy,
	}

	err = d.transactionRepository.CreateTransaction(ctx, tx, entitieTransaction)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return
}

func (d *transactionUsecase) ListTransaction(ctx context.Context) (transactions []domain.TransactionResponse, err error) {
	resp, err := d.transactionRepository.ListTransaction(ctx, d.connection)
	if err != nil {
		return nil, err
	}

	for _, entitie := range resp {
		transaction := domain.TransactionResponse{
			ID:                entitie.ID,
			ContractNumber:    entitie.ContractNumber,
			OTR:               entitie.OTR,
			AdminFee:          entitie.AdminFee,
			InstallmentAmount: entitie.InstallmentAmount,
			AmountInterest:    entitie.AmountInterest,
			ProductName:       entitie.ProductName,
			CreatedBy:         entitie.CreatedBy,
			UpdatedBy:         entitie.UpdatedBy,
			CreatedAt:         entitie.CreatedAt,
			UpdatedAt:         entitie.UpdatedAt,
		}

		transactions = append(transactions, transaction)
	}

	return
}
