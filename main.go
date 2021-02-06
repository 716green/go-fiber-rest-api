package main

import (
	"github.com/gofiber/fiber"
	"github.com/716green/go-fiber-rest-api/tech"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World")
}

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/tech", tech.GetTechnologies)
}

func main() {
	app := fiber.New()

	app.Get("/", helloWorld)

	app.Listen(3000)
}