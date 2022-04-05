package server

import (
	"fmt"

	"github.com/afa7789/flight-min-api/internal/flight"
	fiber "github.com/gofiber/fiber/v2"
)

func FlightPoints(c *fiber.Ctx) error {
	// body struct
	b := struct {
		Paths [][]string `json:"paths"`
	}{}

	// parse the body
	if err := c.BodyParser(&b); err != nil {
		println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := validateEntry(b.Paths); err != nil {
		println(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	answer, err := flight.Parser(b.Paths)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(answer)
}

func validateEntry(entry [][]string) error {
	for _, v := range entry {
		if len(v) != 2 {
			return fmt.Errorf("invalid entry: %v", v)
		}
	}
	return nil
}
