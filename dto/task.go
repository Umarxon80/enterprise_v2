package dto



type InputTask struct {
	CompanyId string `json:"company_id" validate:"required,min=1"`
	TaskField  string `json:"field" validate:"required,min=1"`
	Address     string `json:"address" validate:"required"`
	Investment     string `json:"investment" validate:"required"`
	Document  string `json:"document" validate:"required"`
}

type OutputTask struct {
	Id string `json:"id"`
	User OutputUser `json:"user"`
	Company OutputCompany `json:"company"`
	BusinessPlan OutputBusiness `json:"business_plan_"`
	Node OutputNode `json:"node"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
