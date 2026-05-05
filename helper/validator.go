package helper

import (
	"enterprise_v2/dto"

	"github.com/go-playground/validator/v10"
)

var val *validator.Validate= validator.New(validator.WithRequiredStructEnabled())

type structTypes interface{
	dto.InputCompany | dto.InputRole |dto.InputUser
}

func Validate[T structTypes](st T) error {
	return val.Struct(st)
}