package main

import (
	"enterprise_v2/db"
	"enterprise_v2/handlers"
	customLogger "enterprise_v2/logger"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/earlydata"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/fiber/v3/middleware/responsetime"
	"github.com/lpernett/godotenv"
)

func main() {
	// logger set up
	log.SetOutput(customLogger.Logger())

	// env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loadinng env")
	}
	// caching
	// cacheMiddleware := cache.New(cache.Config{
	// 	Expiration: 10 * time.Second,
	// })


	// db connection
	db.Connect()
	defer db.DbConnection.Close()

	// generating application
	app := fiber.New(fiber.Config{AppName: "Fiber"})
	app.Use(recoverer.New(recoverer.Config{EnableStackTrace: true}))
	app.Use(requestid.New())
	app.Use(responsetime.New())

	// helmet - basic protection
	app.Use(helmet.New())

	// limiter
	app.Use(limiter.New(limiter.Config{
		Max:          20,
		Expiration:   5 * time.Minute,
		KeyGenerator: func(ctx fiber.Ctx) string { return ctx.IP() },
		LimitReached: func(ctx fiber.Ctx) error {
			log.Error("Too many requests user: ", ctx.IP())
			return ctx.Status(429).JSON(fiber.Map{
				"error": "Too many requests try later",
			})
		},
	}))

	// req logs
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} status:${status} - ${method} reqId: ${requestid}, time:${latency}\n",
		Stream: customLogger.Logger(),
		CustomTags: map[string]logger.LogFunc{
			"requestid": func(output logger.Buffer, c fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString(requestid.FromContext(c))
			},
		},
	}))
	// earlydata
	app.Use(earlydata.New())
	// idempotency
	app.Use(idempotency.New(idempotency.Config{Lifetime: 10 * time.Second}))


	companyRouter:=app.Group("/company")
	companyRouter.Get("/",handlers.GetCompanies)
	companyRouter.Get("/:id",handlers.GetOneCompany)
	companyRouter.Post("/",handlers.CreateCompany)
	companyRouter.Patch("/:id",handlers.PatchCompany)
	companyRouter.Delete("/:id",handlers.DeleteCompany)

	roleRouter:=app.Group("/role")
	roleRouter.Get("/",handlers.GetRoles)
	roleRouter.Get("/:id",handlers.GetOneRole)
	roleRouter.Post("/",handlers.CreateRole)
	roleRouter.Patch("/:id",handlers.PatchRole)
	roleRouter.Delete("/:id",handlers.DeleteRole)
	
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}