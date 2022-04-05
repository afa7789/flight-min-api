package server

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestStartEndOfFlight(t *testing.T) {

	tests := []struct {
		name         string
		payload      string
		wantedStatus int
	}{
		{
			name:         "Test with valid payload",
			payload:      `{}`,
			wantedStatus: fiber.StatusOK,
		},
		{
			name:         "Test with invalid payload",
			payload:      `{}`,
			wantedStatus: fiber.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := NewServer()

			req := httptest.NewRequest("GET", "http://google.com", nil)

			resp, _ := s.router.Test(req)

			// Do something with results:
			if resp.StatusCode != tt.wantedStatus {
				t.Errorf("StartEndOfFlight() = %v, want %v", resp.StatusCode, tt.wantedStatus)
			}
		})
	}
}
