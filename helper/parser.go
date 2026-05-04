package helper

import "github.com/gofiber/fiber/v3"

func Parser[T any](ctx fiber.Ctx, st T) (T, error) {
	if err := ctx.Bind().JSON(&st); err != nil {
		return st, err
	}
	return st,nil
}