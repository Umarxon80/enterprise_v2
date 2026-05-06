package dto


type InputBusiness struct {
	CompanyId string `json:"company_id" validate:"required,min=1"`
	BusinessField  string `json:"field" validate:"required,min=1"`
	Address     string `json:"address" validate:"required"`
	Investment     string `json:"investment" validate:"required"`
	Document  string `json:"document" validate:"required"`
}

type OutputBusiness struct {
	Id string `json:"id"`
	Field  string `json:"field" validate:"required,min=1"`
	Address     string `json:"address" validate:"required,email"`
	Document  string `json:"document" validate:"required"`
}
