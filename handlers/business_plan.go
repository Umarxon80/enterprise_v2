package handlers

import (
	"enterprise_v2/db"
	"enterprise_v2/dto"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func UploadBusinessPlan(ctx fiber.Ctx) error {
    var	input_business dto.InputBusiness
	user_id := ctx.Locals("userid").(string)
	user_email := ctx.Locals("email").(string)

	// Parsing
	if err:=ctx.Bind().JSON(&input_business);err!=nil {
		log.Error("Wrong input",err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err":err.Error(),
			"message":"Please fill out form correctly",
		})
	}
	//Validation
	if err:=val.Struct(input_business);err!=nil {
		log.Error("Wrong input",err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err":err.Error(),
			"message":"Please fill out form correctly",
		})
	}

	// Company_user 
	if err:=db.CreateCompanyUser(user_id,input_business.CompanyId);err!=nil {
		log.Error("Error creating company_user",err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err":err.Error(),
			"message":"Please check company_id",
		})
	}

	// Business plan 
	business_id,err:=db.CreateBusinessPlan(input_business)
	if err!=nil {
		log.Error("Error creating business plan",err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err":err.Error(),
		})
	}

	// Task
	task_id,err:=db.CreateTask(user_id,input_business.CompanyId,business_id);
	if err!=nil {
		log.Error("Error creating task",err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err":err.Error(),
		})
	}

	// task_log 
	if err:=db.CreateTaskLog(task_id,1,2,"Task created");err!=nil {
		log.Error("Error creating task_log",err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err":err.Error(),
		})
	}
	
	// Bill 
	if err:=MakeAndSendTaskBill(user_id,user_email);err!=nil {
		log.Error("Error creating bill",err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err":err.Error(),
		})
	}
	log.Info("Business plan was submited id: ",business_id)
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message":"Your business plan was submited",
			"next_actions":"Please pay submission fee. Invoice was sent to your email",
		})
}

func GetReviewerBusunessPlans(ctx fiber.Ctx) error  {
	tasks,err:=db.GetTasks("review")
	if err!=nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return ctx.JSON(tasks)
}