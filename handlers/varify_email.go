package handlers

import (
	"enterprise_v2/db"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)


func VarifyEmail(ctx fiber.Ctx) error {
	code:=ctx.Params("otp")
	
	user_id, err:=db.CheckOtp(code)
	if err!=nil{
		log.Error("Wrong input")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	if err:=db.DeleteOtp(user_id); err!=nil{
		log.Error("Wrong input")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}


	if err:=db.ActivateUser(user_id); err!=nil{
		log.Error("Wrong input")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}


	log.Info("User activated id: ",user_id)
	return  ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":"account activated successfully",
		"id":user_id,
	})
}