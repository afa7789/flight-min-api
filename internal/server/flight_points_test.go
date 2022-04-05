package server

import (
	"net/http/httptest"
	"strings"
	"testing"

	fiber "github.com/gofiber/fiber/v2"
)

func TestFlightPoints(t *testing.T) {
	tests := []struct {
		name         string
		payload      string
		wantedStatus int
	}{
		{
			name:         "Test with valid payload",
			payload:      `{"paths":[["a","b"]]}`,
			wantedStatus: fiber.StatusOK,
		},
		{
			name:         "Test with invalid payload",
			payload:      `erro`,
			wantedStatus: fiber.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServer()

			req := httptest.NewRequest("GET", "/flight_points", strings.NewReader(tt.payload))
			// setting header as json
			req.Header.Set("Content-Type", "application/json; charset=UTF-8")

			resp, _ := s.router.Test(req)

			// Do something with results:
			if resp.StatusCode != tt.wantedStatus {
				t.Errorf("FlightPoints() = %v, want %v", resp.StatusCode, tt.wantedStatus)
			}
		})
	}
}

func Test_validateEntry(t *testing.T) {
	type args struct {
		entry [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid entry",
			args: args{
				entry: [][]string{},
			},
			wantErr: false,
		},
		{
			name: "Test with invalid entry",
			args: args{
				entry: [][]string{
					{"a", "b"},
					{"c", "d", "e"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateEntry(tt.args.entry); (err != nil) != tt.wantErr {
				t.Errorf("validateEntry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
