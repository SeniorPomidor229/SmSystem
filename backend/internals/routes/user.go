package routes

import (
	"sm-system/internals/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(router fiber.Router) {
	user := router.Group("/user")

	user.Post("/", handlers.Register)
}