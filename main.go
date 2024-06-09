package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/tanabebe/myfiber-app/database"
	"github.com/tanabebe/myfiber-app/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	database.Connect()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	router.SetupRoute(app)

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
