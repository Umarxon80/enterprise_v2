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
type Invoice struct{
	Invoice string `json:"invoice" validate:"required"`
}

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
	err = email.SendMail(user_email, "Bob","Application fee", fmt.Sprintf("Please pay %s till %s \n invoice: %s", bill.Amount, bill.Deadline, invoice))
	if err != nil {
		return fmt.Errorf("error sending bill %w", err)
	}
	return nil
}
func MakeAndSendTaskBill(user_Id, user_email string) error {
	bill := dto.InputBill{
		UserId:   user_Id,
		Amount:   "500$",
		Deadline: time.Now().Add(5 * 24 * time.Hour),
	}
	invoice, err := db.CreateBill(bill)
	if err != nil {
		return fmt.Errorf("Error making bill %w", err)
	}
	err = email.SendMail(user_email, "Bob","Task submission fee", fmt.Sprintf("Please pay %s till %s \nInvoice: %s", bill.Amount, bill.Deadline, invoice))
	if err != nil {
		return fmt.Errorf("error sending bill %w", err)
	}
	return nil
}

func PayBill(ctx fiber.Ctx) error {
	var bill Invoice
	id := ctx.Locals("userid").(string)
	if err := ctx.Bind().JSON(&bill); err != nil {
		log.Errorf("Error parsing bill_invoice %w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	if err := db.PayBill(bill.Invoice, id); err != nil {
		log.Errorf("Error paying bill %w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err:= email.SendMail(ctx.Locals("email").(string),"Bob","Profile activated","Your profile was activated. Now you can submit your business plan for review");err!=nil {
		log.Errorf("Error sending email %w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{
		"message":"Apllicatation bill was paid successfully",
		"next_stage":"Now you can proceed to aplying your company info for review",
	})
}
func PayTaskBill(ctx fiber.Ctx) error {
	var bill Invoice
	user_id := ctx.Locals("userid").(string)
	if err := ctx.Bind().JSON(&bill); err != nil {
		log.Errorf("Error parsing bill_invoice %w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	
	// Paying bill
	if err := db.PayBill(bill.Invoice, user_id); err != nil {
		log.Errorf("Error paying bill %w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Getting task
	task, err:=db.GetTaskByUserId(user_id) 
	if err!=nil{
		log.Errorf("Error getting task %w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}


	// Getting Next node
	next_node_id, err:=db.GetNextNode(task.Node.Id) 
	if err!=nil{
		log.Errorf("Error getting next_node %w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if next_node_id==0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "All steps for this task is compleated"})
	}

	// Updating task
	if err := db.TaskUpdateNode(task.Id,next_node_id); err != nil {
		log.Errorf("Error updating task %w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Updating task_log
	if err := db.CreateTaskLog(task.Id,task.Node.Id,next_node_id,"User made paymeny, task is sent to review"); err != nil {
		log.Errorf("Error updating task %w", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// email
	if err:= email.SendMail(ctx.Locals("email").(string),task.User.FirstName+" "+task.User.LastName,"Business plan application fee","Your application fee was Paid. Every update on your application will be sent to this email. Please stay in touch");err!=nil {
		log.Errorf("Error sending email %w", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{
		"message":"Apllicatation bill was paid successfully",
		"next_stage":"Please wait for updates",
	})
}
