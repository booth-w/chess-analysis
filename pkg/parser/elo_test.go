package parser

import (
	"testing"
)

func TestParseElo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		hasErr   bool
	}{
		{
			name:     "Valid Elo",
			input:    `[WhiteElo "1500"]`,
			expected: 1500,
			hasErr:   false,
		},
		{
			name:     "Unknown Elo",
			input:    `[WhiteElo "?"]`,
			expected: -1,
			hasErr:   false,
		},
		{
			name:     "Invalid Elo",
			input:    `[WhiteElo "invalid"]`,
			expected: -1,
			hasErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := parseElo(tt.input)
			if (err != nil) != tt.hasErr {
				t.Errorf("Expected error: %v, got: %v", tt.hasErr, err)
			}
			if output != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, output)
			}
		})
	}
}

func TestParseEloFilter(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedMin int
		expectedMax int
		hasErr      bool
	}{
		{
			name:        "Empty",
			input:       "",
			expectedMin: 0,
			expectedMax: 9999,
			hasErr:      false,
		},
		{
			name:        "Min 1500",
			input:       "1500",
			expectedMin: 1500,
			expectedMax: 9999,
			hasErr:      false,
		},
		{
			name:        "Min 1500 Max 2000",
			input:       "1500-2000",
			expectedMin: 1500,
			expectedMax: 2000,
			hasErr:      false,
		},
		{
			name:        "Invalid format",
			input:       "invalid",
			expectedMin: -1,
			expectedMax: -1,
			hasErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outMin, outMax, err := ParseEloFilter(tt.input)
			if (err != nil) != tt.hasErr {
				t.Errorf("Expected hasErr=%v, got %v", tt.hasErr, err)
			}
			if outMin != tt.expectedMin {
				t.Errorf("Expected %v, got %v", tt.expectedMin, outMin)
			}
			if outMax != tt.expectedMax {
				t.Errorf("Expected %v, got %v", tt.expectedMax, outMax)
			}
		})
	}
}
