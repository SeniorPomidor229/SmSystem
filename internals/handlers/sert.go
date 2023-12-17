package handlers

import (
	"sm-system/database"
	"sm-system/internals/models"

	"github.com/gofiber/fiber/v2"
)


func GetSerts(c *fiber.Ctx) error {
	db := database.DB
	var serts []models.Sert

	db.Find(&serts)

	if len(serts) <= 0 {
		return c.Status(404).JSON(fiber.Map{"error":"serts not found"})
	}

	return c.JSON(fiber.Map{"detail":serts})
}