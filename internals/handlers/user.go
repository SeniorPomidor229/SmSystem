package handlers

import (
	"sm-system/database"
	"sm-system/internals/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Register(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.User)

	err := c.BodyParser(&user); if err != nil {
		return c.Status(500).JSON(fiber.Map{"error":err.Error()})
	}

	user.ID = uuid.New()

	err = db.Create(&user).Error; if err != nil {
		return c.Status(500).JSON(fiber.Map{"error":err.Error()})
	}

	return c.JSON(fiber.Map{"detail":"user successfull create"})
}