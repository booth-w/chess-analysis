package parser

import (
	"testing"
)

func TestParseEvent(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected string
		hasErr   bool
	}{
		{
			name:     "Event with link",
			line:     `[Event "Rated Bullet tournament https://lichess.org/tournament/yc1WW2Ox"]`,
			expected: "Rated Bullet tournament",
			hasErr:   false,
		},
		{
			name:     "Event without link",
			line:     `[Event "Rated Classical game"]`,
			expected: "Rated Classical game",
			hasErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := parseEvent(tt.line)
			if (err != nil) != tt.hasErr {
				t.Errorf("Expected error: %v, got: %v", tt.hasErr, err)
			}
			if output != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, output)
			}
		})
	}

}
