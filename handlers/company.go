package handlers

import (
	"enterprise_v2/db"
	"enterprise_v2/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)
var val *validator.Validate = validator.New(validator.WithRequiredStructEnabled())
func GetCompanies(ctx fiber.Ctx) error {
	companies, err := db.GetCompanies()
	if err != nil {
		log.Errorf("Error getting Companies, err:", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(companies)
}
func GetOneCompany(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	company, err := db.GetOneCompany(id)
	if err != nil {
		log.Errorf("Error getting Company, err:", err)
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(company)
}
func CreateCompany(ctx fiber.Ctx) error {
	var company dto.InputCompany
	if err:=ctx.Bind().JSON(&company);err!=nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := val.Struct(company)
	if err != nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}


	id, err := db.CreateCompany(company)

	if err != nil {
		log.Errorf("Error creating Company, err:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"company created, id:": id})
}
func PatchCompany(ctx fiber.Ctx) error {
	var company dto.InputCompany
	id := ctx.Params("id")
	if err:=ctx.Bind().JSON(&company);err!=nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := val.Struct(company)
	if err != nil {
		log.Error("Wrong input ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	
	id, err = db.PatchCompany(id, company)
	if err != nil {
		log.Errorf("Error updating Company, err:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"company updated, id:": id})
}
func DeleteCompany(ctx fiber.Ctx) error {
	id := ctx.Params("id")

	id, err := db.DeleteCompany(id)
	if err != nil {
		log.Errorf("Error deleting Company, err:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"company deleted, id:": id})
}
