package helper

import 	"github.com/go-playground/validator/v10"

var val *validator.Validate= validator.New(validator.WithRequiredStructEnabled())


func Validate[T any](st T) error {
	return val.Struct(st)
}