package parser

import (
	"testing"
)

func TestGetWinner(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		hasErr   bool
	}{
		{
			name:     "White wins",
			input:    `[Result "1-0"]`,
			expected: 0,
			hasErr:   false,
		},
		{
			name:     "Black wins",
			input:    `[Result "0-1"]`,
			expected: 1,
			hasErr:   false,
		},
		{
			name:     "Draw",
			input:    `[Result "1/2-1/2"]`,
			expected: 2,
			hasErr:   false,
		},
		{
			name:     "Invalid game",
			input:    `[Result "*"]`,
			expected: 3,
			hasErr:   false,
		},
		{
			name:     "Invalid input",
			input:    `[Result "invalid"]`,
			expected: -1,
			hasErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := getWinner(tt.input)
			if (err != nil) != tt.hasErr {
				t.Errorf("Expected error: %v, got: %v", tt.hasErr, err)
			}
			if output != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, output)
			}
		})
	}
}
