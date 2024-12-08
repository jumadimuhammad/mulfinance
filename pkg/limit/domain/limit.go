package domain

import (
	"time"
)

type Limit struct {
	ID        uint      `json:"id"`
	Limit     float64   `json:"limit"`
	Tenor     uint      `json:"tenor"`
	CreatedBy uint      `json:"created_by"`
	UpdatedBy uint      `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
