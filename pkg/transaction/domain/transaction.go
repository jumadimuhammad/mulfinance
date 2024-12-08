package domain

import (
	"time"
)

type Transaction struct {
	ID                uint      `json:"id"`
	DebiturID         uint      `json:"debitur_id"`
	DebiturLimitID    uint      `json:"debitur_limit_id"`
	OTR               float64   `json:"otr"`
	AdminFee          float64   `json:"admin_fee"`
	InstallmentAmount float64   `json:"installment_amount"`
	AmountInterest    float64   `json:"amount_interest"`
	ProductName       string    `json:"product_name"`
	CreatedBy         uint      `json:"created_by"`
	UpdatedBy         uint      `json:"updated_by"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type TransactionResponse struct {
	ID                uint      `json:"id"`
	ContractNumber    string    `json:"contract_number"`
	OTR               float64   `json:"otr"`
	AdminFee          float64   `json:"admin_fee"`
	InstallmentAmount float64   `json:"installment_amount"`
	AmountInterest    float64   `json:"amount_interest"`
	ProductName       string    `json:"product_name"`
	CreatedBy         uint      `json:"created_by"`
	UpdatedBy         uint      `json:"updated_by"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
