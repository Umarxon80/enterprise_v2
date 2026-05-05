package dto

type InputRole struct {
	Name             string `json:"name" validate:"required,min=2"`
}

type OutputRole  struct {
	Id                string  `json:"id"`
	Name              string `json:"name"`
}
