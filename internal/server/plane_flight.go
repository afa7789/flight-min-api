package server

import (
	fiber "github.com/gofiber/fiber/v2"
)

func StartEndOfFlight(c *fiber.Ctx) error {
	// body struct
	b := struct {
		Paths [][2]string `json:"paths"`
	}{}

	// parse the body
	if err := c.BodyParser(&b); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(true)
}
