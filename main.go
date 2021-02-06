package main

import (
	"fmt"
	"github.com/716green/go-fiber-api/database"
	"github.com/716green/go-fiber-api/tech"
	"github.com/gofiber/fiber"

	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/tech", tech.GetTechnologies)
	app.Get("/api/v1/tech/:id", tech.GetTechnology)
	app.Post("/api/v1/tech", tech.NewTechnology)
	app.Delete("/api/v1/tech/:id", tech.DeleteTechnology)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "technologies.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	database.DBConn.AutoMigrate(&tech.Tech{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}