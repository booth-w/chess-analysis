package parser

import (
	"strings"
)

// Removes the lichess event link
// Example: [Event "Rated Bullet tournament https://lichess.org/tournament/yc1WW2Ox"] -> Rated Bullet tournament
func parseEvent(line string) (string, error) {
	line, err := parseGeneric(line)
	if err != nil {
		return "", err
	}

	index := strings.Index(line, ":")
	if index != -1 {
		line = line[:index-6]
	}

	return line, nil
}
