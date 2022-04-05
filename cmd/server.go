package cmd

import (
	"github.com/afa7789/flight-min-api/internal/domain"
	"github.com/afa7789/flight-min-api/internal/server"
)

// execute the server package to start the server running
func ServerExecute(f domain.Flags) error {
	s := server.NewServer()
	s.Start(*f.Port)
	return nil
}
