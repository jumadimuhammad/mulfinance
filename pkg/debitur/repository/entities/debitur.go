package entities

import (
	"mulfinance/pkg/debitur/domain"
	"time"
)

type Debitur struct {
	ID          uint      `gorm:"primaryKey"`
	NIK         string    `gorm:"type:varchar(16);not null;unique"`
	FullName    string    `gorm:"type:varchar(100);not null"`
	LegalName   string    `gorm:"type:varchar(100);not null"`
	BirthPlace  string    `gorm:"type:varchar(100);not null"`
	BirthDate   time.Time `gorm:"type:datetime;not null"`
	Salary      float64   `gorm:"type:double(10,2);not null;default:0"`
	KTPPhoto    string    `gorm:"type:varchar(255);not null"`
	SelfiePhoto string    `gorm:"type:varchar(255);not null"`
	CreatedBy   uint      `gorm:"type:int;not null"`
	UpdatedBy   uint      `gorm:"type:int;not null"`
	CreatedAt   time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

func (Debitur) TableName() string {
	return "debiturs"
}

func (d *Debitur) ToDomain() *domain.Debitur {
	return &domain.Debitur{
		ID:          d.ID,
		NIK:         d.NIK,
		FullName:    d.FullName,
		LegalName:   d.LegalName,
		BirthPlace:  d.BirthPlace,
		BirthDate:   d.BirthDate,
		Salary:      d.Salary,
		KTPPhoto:    d.KTPPhoto,
		SelfiePhoto: d.SelfiePhoto,
		CreatedBy:   d.CreatedBy,
		UpdatedBy:   d.UpdatedBy,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}

}
