package main

import (
	"github.com/samiklah/gocrud/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)
	main1()

	app.Listen(":3000")
}
