package handlers

import (
	"enterprise_v2/db"
	"enterprise_v2/dto"
	"enterprise_v2/helper"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func GetRoles(ctx fiber.Ctx) error {
	roles, err := db.GetRoles()
	if err != nil {
		log.Errorf("Error getting Roles, err:", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(roles)
}
func GetOneRole(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	role, err := db.GetOneRole(id)
	if err != nil {
		log.Errorf("Error getting Role, err:", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(role)
}
func CreateRole(ctx fiber.Ctx) error {
	var role dto.InputRole
	if err:=ctx.Bind().JSON(&role);err!=nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := helper.Validate(role)
	if err != nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}


	id, err := db.CreateRole(role)

	if err != nil {
		log.Errorf("Error creating Role, err:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"role created, id:": id})
}
func PatchRole(ctx fiber.Ctx) error {
	var role dto.InputRole
	id := ctx.Params("id")
	if err:=ctx.Bind().JSON(&role);err!=nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := helper.Validate(role)
	if err != nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	
	id, err = db.PatchRole(id, role)
	if err != nil {
		log.Errorf("Error updating Role, err:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"role updated, id:": id})
}
func DeleteRole(ctx fiber.Ctx) error {
	id := ctx.Params("id")

	id, err := db.DeleteRole(id)
	if err != nil {
		log.Errorf("Error deleting Role, err:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"role deleted, id:": id})
}
