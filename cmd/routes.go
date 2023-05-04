package main

import (
	"github.com/samiklah/gocrud/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/facts", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
