package parser

import (
	"testing"
)

func TestGetWinner(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "White wins",
			input:    "1-0",
			expected: 0,
		},
		{
			name:     "Black wins",
			input:    "0-1",
			expected: 1,
		},
		{
			name:     "Draw",
			input:    "1/2-1/2",
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := GetWinner(tt.input)
			if output != tt.expected {
				t.Errorf("got %v, expected %v", output, tt.expected)
			}
		})
	}
}
