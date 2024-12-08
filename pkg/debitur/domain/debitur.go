package domain

import (
	limitDomain "mulfinance/pkg/limit/domain"
	transactionDomain "mulfinance/pkg/transaction/domain"
	"time"
)

type Debitur struct {
	ID          uint      `json:"id"`
	NIK         string    `json:"nik"`
	FullName    string    `json:"full_name"`
	LegalName   string    `json:"legal_name"`
	BirthPlace  string    `json:"birth_place"`
	BirthDate   time.Time `json:"birth_date"`
	Salary      float64   `json:"salary"`
	KTPPhoto    string    `json:"ktp_photo"`
	SelfiePhoto string    `json:"selfie_photo"`
	CreatedBy   uint      `json:"created_by"`
	UpdatedBy   uint      `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DebiturResponse struct {
	ID          uint      `json:"id"`
	NIK         string    `json:"nik"`
	FullName    string    `json:"full_name"`
	LegalName   string    `json:"legal_name"`
	BirthPlace  string    `json:"birth_place"`
	BirthDate   time.Time `json:"birth_date"`
	Salary      float64   `json:"salary"`
	KTPPhoto    string    `json:"ktp_photo"`
	SelfiePhoto string    `json:"selfie_photo"`
	CreatedBy   uint      `json:"created_by"`
	UpdatedBy   uint      `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DebiturLimitResponse struct {
	Limit       []limitDomain.Limit                     `json:"limits"`
	Transaction []transactionDomain.TransactionResponse `json:"transactions"`
}

type DebiturTransactionResponse struct {
	Transaction []transactionDomain.TransactionResponse `json:"transactions"`
}

type DebiturTransaction struct {
	ID          uint                            `json:"id"`
	NIK         string                          `json:"nik"`
	FullName    string                          `json:"full_name"`
	LegalName   string                          `json:"legal_name"`
	BirthPlace  string                          `json:"birth_place"`
	BirthDate   time.Time                       `json:"birth_date"`
	Salary      float64                         `json:"salary"`
	KTPPhoto    string                          `json:"ktp_photo"`
	SelfiePhoto string                          `json:"selfie_photo"`
	Transaction []transactionDomain.Transaction `json:"transactions"`
	CreatedBy   uint                            `json:"created_by"`
	UpdatedBy   uint                            `json:"updated_by"`
	CreatedAt   time.Time                       `json:"created_at"`
	UpdatedAt   time.Time                       `json:"updated_at"`
}
