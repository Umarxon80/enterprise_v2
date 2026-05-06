package handlers

import (
	"enterprise_v2/db"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)
type body_code struct {
	Code string `json:"code" validate:"min=6"`
}
func VarifyEmail(ctx fiber.Ctx) error {
	user_id:=ctx.Locals("userid").(string)
	user_email:=ctx.Locals("email").(string)
	var code body_code
	if err:=ctx.Bind().JSON(&code); err!=nil{
		log.Error("Wrong input")
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	
	if err:=db.CheckOtp(user_id,code.Code); err!=nil{
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


	if err:=MakeAndSendBill(user_id,user_email); err!=nil {
		log.Error("Error making and sending bill err:" ,err.Error())
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	 
	log.Info("User activated id: ",user_id)
	return  ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":"account activated successfully",
		"next_actions":"please pay application fee. Invoice was sent to your account",
		"id":user_id,
	})
}