package handlers

import (
	"time"

	"sm-system/database"
	"sm-system/internals/models"

	"github.com/gofiber/fiber/v2"
)


func GetSerts(c *fiber.Ctx) error {
	db := database.DB
	var serts []models.Sert

	queryDate := c.Params("date") // Получение параметра даты из URL

	if queryDate != "" {
		parsedDate, err := time.Parse("02.01.2006", queryDate)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid date format. Please use format DD.MM.YYYY"})
		}
		db.Where("date_validity = ?", parsedDate).Find(&serts)
	} else {
		db.Find(&serts)
	}

	if len(serts) <= 0 {
		return c.Status(404).JSON(fiber.Map{"error": "serts not found"})
	}

	return c.JSON(fiber.Map{"detail": serts})
}