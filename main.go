package main

import (
	"log"
	"time"

	"sm-system/database"
	"sm-system/internals/workers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Pivo"})
	})

	database.ConnectDB()

	go func() {
		for {
			err := workers.FetchDataAndSave(database.DB); if err != nil {
				log.Println("Error", err)
			}
			time.Sleep(1 * time.Minute)
		}
	}()


	app.Listen(":8000")
}