package tech

import (
	"github.com/716green/go-fiber-api/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Tech struct {
	gorm.Model
	Technology string `json:"technology"`
	Category string `json:"category"`
	Note string `json:"note"`
}

func GetTechnologies(c *fiber.Ctx) {
	// curl http://localhost:3000/api/v1/tech

	db := database.DBConn
	var tech []Tech
	db.Find(&tech)
	c.JSON(tech)
}

func GetTechnology(c *fiber.Ctx) {
	// curl http://localhost:3000/api/v1/tech/1

	id := c.Params("id")
	db := database.DBConn
	var tech Tech
	db.Find(&tech, id)
	c.JSON(tech)
}

func NewTechnology(c *fiber.Ctx) {
	db := database.DBConn
	// curl -X POST -H "Content-Type: application/json" --data "{\"technology\":\"Go\", \"category\":\"System Level Language\", \"note\": \"Server Programming Language by Google\"}" http://localhost:3000/api/v1/tech

	tech := new(Tech)
	if err := c.BodyParser(tech); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&tech)
	c.JSON(tech)
}

func DeleteTechnology(c *fiber.Ctx) {
	// curl -X DELETE  http://localhost:3000/api/v1/tech/1
	id := c.Params("id")
	db := database.DBConn
	var tech Tech
	db.First(&tech, id)
	if tech.Technology == "" {
		c.Status(500).Send("No technology found with this name")
		return
	}
	db.Delete(&tech)
	c.Send("Technology Successfully Deleted")
}