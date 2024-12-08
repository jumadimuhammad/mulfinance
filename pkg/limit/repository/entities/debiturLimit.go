package entities

import (
	"time"
)

type DebiturLimit struct {
	ID        uint      `gorm:"primaryKey"`
	DebiturID uint      `gorm:"type:int;not null"`
	LimitID   uint      `gorm:"type:int;not null"`
	CreatedBy uint      `gorm:"type:int;not null"`
	UpdatedBy uint      `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (DebiturLimit) TableName() string {
	return "debitur_limits"
}
