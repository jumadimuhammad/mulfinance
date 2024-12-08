package entities

import (
	"time"
)

type Transaction struct {
	ID                 uint      `gorm:"primaryKey"`
	TransactionLimitID uint      `gorm:"type:int;not null"`
	ContractNumber     string    `gorm:"type:varchar(10);not null"`
	OTR                float64   `gorm:"type:double(10,2);not null;default:0"`
	AdminFee           float64   `gorm:"type:double(10,2);not null;default:0"`
	InstallmentAmount  float64   `gorm:"type:double(10,2);not null;default:0"`
	AmountInterest     float64   `gorm:"type:double(10,2);not null;default:0"`
	ProductName        string    `gorm:"column:product_name"`
	CreatedBy          uint      `gorm:"type:int;not null"`
	UpdatedBy          uint      `gorm:"type:int;not null"`
	CreatedAt          time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (Transaction) TableName() string {
	return "transactions"
}
