package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tanabebe/myfiber-app/handler"
)

func SetupRoute(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", handler.GetAllUsers)
}
