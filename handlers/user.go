package handlers

import (
	"enterprise_v2/db"
	"enterprise_v2/dto"
	"enterprise_v2/helper"
	"enterprise_v2/otp"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func GetUsers(ctx fiber.Ctx) error {
	users, err := db.GetUsers()
	if err != nil {
		log.Errorf("Error getting Users, err:", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(users)
}
func GetOneUser(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := db.GetOneUser(id)
	if err != nil {
		log.Errorf("Error getting User, err:", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(user)
}
func CreateUser(ctx fiber.Ctx) error {
	var user dto.InputUser
	// Body parsing
	if err:=ctx.Bind().JSON(&user);err!=nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validation
	err := val.Struct(user)
	if err != nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Hashing password
	user.Password,err = helper.HashPassword(user.Password)
	if err != nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}


	id, err := db.CreateUser(user)
	if err!=nil {
		log.Errorf("Error creating user %v",err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err = otp.MakeAndSendOtp(id, user.Email); err != nil {
		log.Errorf("Error sending otp %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"user created, id:": id})
}
func PatchUser(ctx fiber.Ctx) error {
	var user dto.InputUser
	id := ctx.Params("id")
	
	// Body parsing
	if err:=ctx.Bind().JSON(&user);err!=nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validation
	err := val.Struct(user)
	if err != nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Hashing password
	user.Password,err = helper.HashPassword(user.Password)
	if err != nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	
	id, err = db.PatchUser(id, user)
	if err != nil {
		log.Errorf("Error updating User, err:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"user updated, id:": id})
}
func DeleteUser(ctx fiber.Ctx) error {
	id := ctx.Params("id")

	id, err := db.DeleteUser(id)
	if err != nil {
		log.Errorf("Error deleting User, err:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"user deleted, id:": id})
}
