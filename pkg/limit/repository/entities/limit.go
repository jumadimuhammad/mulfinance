package entities

import (
	"time"
)

type Limit struct {
	ID        uint      `gorm:"primaryKey"`
	Limit     float64   `gorm:"type:double(10,2);not null;default:0"`
	Tenor     uint      `gorm:"type:int;not null"`
	CreatedBy uint      `gorm:"type:int;not null"`
	UpdatedBy uint      `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (Limit) TableName() string {
	return "limits"
}
