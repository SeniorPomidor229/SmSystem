package routes

import (
	"sm-system/internals/handlers"

	"github.com/gofiber/fiber/v2"
)


func SetupSertRoute(router fiber.Router) {
	sert := router.Group("/sert")

	sert.Get("/", handlers.GetSerts)
}