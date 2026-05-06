package handlers

import (
	"enterprise_v2/db"

	"github.com/gofiber/fiber/v3"
)

func Test(ctx fiber.Ctx) error {
	user_id:=ctx.Locals("userid").(string)
	task,err:=db.GetTaskByUserId(user_id)
	if err!=nil {
		return ctx.JSON(err.Error())
	}
	return ctx.JSON(task)
}