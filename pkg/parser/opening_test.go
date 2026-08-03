package parser

import (
	"reflect"
	"testing"

	"github.com/booth-w/chess-analysis/pkg/game"
)

func TestParseOpening(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected game.Opening
		hasErr   bool
	}{
		{
			name: "Opening with family and variation",
			line: `[Opening "Sicilian Defense: Najdorf Variation, English Attack"]`,
			expected: game.Opening{
				Family:    "Sicilian Defense",
				Variation: []string{"Najdorf Variation", "English Attack"},
			},
			hasErr: false,
		},
		{
			name: "Opening with only family",
			line: `[Opening "Sicilian Defense"]`,
			expected: game.Opening{
				Family:    "Sicilian Defense",
				Variation: nil,
			},
			hasErr: false,
		},
		{
			name:     "Invalid opening format",
			line:     `[Opening "Sicilian Defense: Najdorf Variation: English Attack"]`,
			expected: game.Opening{},
			hasErr:   true,
		},
		{
			name:     "Empty opening",
			line:     `[Opening ""]`,
			expected: game.Opening{},
			hasErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := parseOpening(tt.line)
			if (err != nil) != tt.hasErr {
				t.Errorf("Expected error: %v, got: %v", tt.hasErr, err)
			}
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, output)
			}
		})
	}
}
