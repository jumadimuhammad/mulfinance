package dto

type CreateTransaction struct {
	DebiturID         uint    `json:"debitur_id"`
	DebiturLimitID    uint    `json:"debitur_limit_id"`
	OTR               float64 `json:"otr"`
	AdminFee          float64 `json:"admin_fee"`
	InstallmentAmount float64 `json:"installment_amount"`
	AmountInterest    float64 `json:"amount_interest"`
	ProductName       string  `json:"product_name"`
}
