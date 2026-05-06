package dto

type InputCompany struct {
	Name             string `json:"name" validate:"required,min=2"`
	FinancialHistory string `json:"financial_history" validate:"required"`
	Industry         string `json:"industry" validate:"required,min=8"`
}

type OutputCompany struct {
	Id                string   `json:"id"`
	Name              string `json:"name"`
	FinancialHistory string `json:"financial_history" `
	Industry          string `json:"industry"`
}
