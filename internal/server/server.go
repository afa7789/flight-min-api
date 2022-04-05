package server

import (
	"fmt"
	"log"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
)

// Server is the definition of a REST server based on Gin
type Server struct {
	router *fiber.App
}

// New returns a new server that takes advantage of zerolog for logging
// and holds a reference to the app configuration
func NewServer() *Server {
	r := fiber.New()

	r.Get("/flight_points", FlightPoints)

	return &Server{
		router: r,
	}
}

// Start starts the REST server
// simple mocking pattern to switch port if one is already in use.
func (s *Server) Start(port int) {
	err := s.router.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		// Using this error treatment to try again on next port
		if strings.Contains(err.Error(), "address already in use") {
			fmt.Println("")
			log.Printf("PORT ALREADY IN USE::%d", port)
			port++
			log.Printf("TRYING NEXT PORT:%d\n", port)
			s.Start(port)
		} else {
			panic(err)
		}
	}
}
