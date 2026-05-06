package handlers

import (
	"enterprise_v2/db"
	"enterprise_v2/dto"
	"enterprise_v2/email"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func MakeAndSendBill(user_Id, user_email string) error {
	bill := dto.InputBill{
		UserId:   user_Id,
		Amount:   "50$",
		Deadline: time.Now().Add(5 * 24 * time.Hour),
	}
	invoice, err := db.CreateBill(bill)
	if err != nil {
		return fmt.Errorf("Error making bill %w", err)
	}
	err = email.SendMail(user_email, "Application fee", fmt.Sprintf("Please pay %s till %s \n invoice: %s", bill.Amount, bill.Deadline, invoice))
	if err != nil {
		return fmt.Errorf("error sending bill %w", err)
	}
	return nil
}

func PayBill(ctx fiber.Ctx) error {
	var bill dto.OutputBill
	id := ctx.Locals("userid").(string)
	if err := ctx.Bind().JSON(&bill); err != nil {
		log.Errorf("Error parsing bill_invoice %w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	if err := db.PayBill(bill.Invoice, id); err != nil {
		log.Errorf("Error paying bill %w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{
		"message":"Apllicatation bill was paid successfully",
		"next_stage":"Now you can proceed to aplying your company info for review",
	})
}
