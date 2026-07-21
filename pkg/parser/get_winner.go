package parser

import (
	"fmt"
)

// Expected inputs:
//
//	[Result "1-0"]
//	[Result "0-1"]
//	[Result "1/2-1/2"]
//
// Returns:
//
//	0: white win
//	1: black win
//	2: draw
func getWinner(line string) (int, error) {
	line, err := parseGeneric(line)
	if err != nil {
		return -1, err
	}

	if line == "1/2-1/2" {
		return 2, nil // draw
	} else if line[0] == '1' {
		return 0, nil // white
	} else if line[0] == '0' {
		return 1, nil // black
	} else if line[0] == '*' {
		return 3, nil // invalid
	} else {
		return -1, fmt.Errorf("invalid result %q", line)
	}
}
