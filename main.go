package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vendz/custom-0auth/database"
	"github.com/vendz/custom-0auth/helper"
	"github.com/vendz/custom-0auth/routes"
)

func main() {
	helper.LoadEnv()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	handler := database.NewDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API is Up and Running...")
	})
	routes.UserRoutes(app, &handler)
	routes.ClientRoutes(app, &handler)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
