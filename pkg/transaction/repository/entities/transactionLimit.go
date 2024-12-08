package entities

import (
	"time"
)

type TransactionLimit struct {
	ID             uint      `gorm:"primaryKey"`
	DebiturID      uint      `gorm:"type:int;not null"`
	DebiturLimitID uint      `gorm:"type:int;not null"`
	Status         string    `gorm:"type:enum('completed','processed');not null;default:'processed'"`
	CreatedBy      uint      `gorm:"type:int;not null"`
	UpdatedBy      uint      `gorm:"type:int;not null"`
	CreatedAt      time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (TransactionLimit) TableName() string {
	return "transaction_limits"
}

type TransactionLimitTotal struct {
	Limit       float64 `gorm:"column:credit_limit"`
	TotalCredit float64 `gorm:"column:credit_total"`
}
