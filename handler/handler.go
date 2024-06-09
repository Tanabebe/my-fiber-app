package handler

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/tanabebe/myfiber-app/database"
	"github.com/tanabebe/myfiber-app/model"
)

func GetAllUsers(c fiber.Ctx) error {
	ctx := context.Background()

	db := database.DB.Db

	var users []model.User

	err := db.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		return err
	}
	fmt.Println(c.Status(fiber.StatusOK).JSON(users))
	return c.Status(fiber.StatusOK).JSON(users)
}
