package usecase

import (
	"context"
	"errors"
	"mulfinance/pkg/debitur/domain"
	"mulfinance/pkg/debitur/ports"
	limitDomain "mulfinance/pkg/limit/domain"
	limitPorts "mulfinance/pkg/limit/ports"
	transactionDomain "mulfinance/pkg/transaction/domain"
	transactionPorts "mulfinance/pkg/transaction/ports"

	"gorm.io/gorm"
)

type debiturUsecase struct {
	debiturRepository     ports.IRepository
	limitRepository       limitPorts.IRepository
	transactionRepository transactionPorts.IRepository
	connection            *gorm.DB
}

func NewUsecase(debRepo ports.IRepository, limRepo limitPorts.IRepository, traRepo transactionPorts.IRepository, connection *gorm.DB) ports.IUsecase {
	return &debiturUsecase{debRepo, limRepo, traRepo, connection}
}

func (d *debiturUsecase) ListDebitur(ctx context.Context) (debiturs []domain.DebiturResponse, err error) {
	resp, err := d.debiturRepository.ListDebitur(ctx)
	if err != nil {
		return nil, err
	}

	for _, debitur := range resp {
		debiturs = append(debiturs, domain.DebiturResponse{
			ID:          debitur.ID,
			NIK:         debitur.NIK,
			FullName:    debitur.FullName,
			LegalName:   debitur.LegalName,
			BirthPlace:  debitur.BirthPlace,
			BirthDate:   debitur.BirthDate,
			Salary:      debitur.Salary,
			KTPPhoto:    debitur.KTPPhoto,
			SelfiePhoto: debitur.SelfiePhoto,
			CreatedBy:   debitur.CreatedBy,
			UpdatedBy:   debitur.UpdatedBy,
			CreatedAt:   debitur.CreatedAt,
			UpdatedAt:   debitur.UpdatedAt,
		})
	}

	return
}

func (d *debiturUsecase) GetDebiturByID(ctx context.Context, debiturID uint) (debitur *domain.DebiturResponse, err error) {
	resp, err := d.debiturRepository.GetDebiturByID(ctx, debiturID)
	if err != nil {
		return
	}

	debitur = &domain.DebiturResponse{
		ID:          resp.ID,
		NIK:         resp.NIK,
		FullName:    resp.FullName,
		LegalName:   resp.LegalName,
		BirthPlace:  resp.BirthPlace,
		BirthDate:   resp.BirthDate,
		Salary:      resp.Salary,
		KTPPhoto:    resp.KTPPhoto,
		SelfiePhoto: resp.SelfiePhoto,
		CreatedBy:   resp.CreatedBy,
		UpdatedBy:   resp.UpdatedBy,
		CreatedAt:   resp.CreatedAt,
		UpdatedAt:   resp.UpdatedAt,
	}

	return
}

func (d *debiturUsecase) GetDebiturLimit(ctx context.Context, debiturID uint) (debitur domain.DebiturLimitResponse, err error) {
	respTransLimit, err := d.transactionRepository.GetTransactionLimitProcessedByDebiturID(ctx, d.connection, debiturID)
	if err != nil {
		return
	}

	limitIDs := []uint{}
	if respTransLimit.DebiturLimitID == 0 {
		respDebLimit, err := d.debiturRepository.GetDebiturLimit(ctx, debiturID)
		if err != nil {
			return debitur, err
		}

		for _, limit := range respDebLimit {
			limitIDs = append(limitIDs, limit.LimitID)
		}

	} else {
		respDebLimitID, err := d.debiturRepository.GetDebiturLimitByID(ctx, respTransLimit.DebiturLimitID)
		if err != nil {
			return debitur, err
		}

		limitIDs = append(limitIDs, respDebLimitID.LimitID)
	}

	if len(limitIDs) == 0 {
		return debitur, errors.New("debitur don't have limit")
	}

	respLimit, err := d.limitRepository.GetLimitByIDs(ctx, limitIDs)
	if err != nil {
		return
	}

	for _, entitie := range respLimit {
		limit := limitDomain.Limit{
			ID:        entitie.ID,
			Limit:     entitie.Limit,
			Tenor:     entitie.Tenor,
			CreatedBy: entitie.CreatedBy,
			UpdatedBy: entitie.UpdatedBy,
			CreatedAt: entitie.CreatedAt,
			UpdatedAt: entitie.UpdatedAt,
		}

		debitur.Limit = append(debitur.Limit, limit)
	}

	resp, err := d.transactionRepository.GetTransactionByDebiturID(ctx, d.connection, debiturID)
	if err != nil {
		return
	}

	for _, entitie := range resp {
		transaction := transactionDomain.TransactionResponse{
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

		debitur.Transaction = append(debitur.Transaction, transaction)
	}

	return
}

func (d *debiturUsecase) GetDebiturTransaction(ctx context.Context, debiturID uint) (transactions []transactionDomain.TransactionResponse, err error) {
	resp, err := d.transactionRepository.GetTransactionByDebiturID(ctx, d.connection, debiturID)
	if err != nil {
		return
	}

	for _, entitie := range resp {
		transaction := transactionDomain.TransactionResponse{
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
