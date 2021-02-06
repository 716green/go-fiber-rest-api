package tech

import(
	"github.com/gofiber/fiber"
)

func GetTechnologies(c *fiber.Ctx) {
	c.Send("All Technologies")
}

func GetTechnology(c *fiber.Ctx) {
	c.Send("Single Technology")
}

func NewTechnology(c *fiber.Ctx) {
	c.Send("Add a new technology")
}

func DeleteTechnology(c *fiber.Ctx) {
	c.Send("Delete a technology")
}