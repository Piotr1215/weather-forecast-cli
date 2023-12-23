package main

import (
	"testing"
)

func TestCalculateDaylightDuration(t *testing.T) {
	tests := []struct {
		sunrise string
		sunset  string
		hours   int
		minutes int
	}{
		{"06:00 AM", "06:00 PM", 12, 0},
		{"06:00 AM", "07:30 PM", 13, 30},
		{"06:00 AM", "05:30 PM", 11, 30},
		{"06:00 AM", "06:30 PM", 12, 30},
	}

	for _, tt := range tests {
		hours, minutes, err := calculateDaylightDuration(tt.sunrise, tt.sunset)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if hours != tt.hours || minutes != tt.minutes {
			t.Errorf("incorrect result for sunrise=%s, sunset=%s: got %d hours and %d minutes, expected %d hours and %d minutes",
				tt.sunrise, tt.sunset, hours, minutes, tt.hours, tt.minutes)
		}
	}
}
