package main

import (
	"github.com/divrhino/divrhino-trivia/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
