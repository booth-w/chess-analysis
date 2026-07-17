package parser

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

// Converts a PGN elo string to an int. Returns -1 if the elo is unknown ("?").
func parseElo(line string) int {
	eloStr := parseGeneric(line)
	if eloStr == "?" {
		return -1
	}

	elo, err := strconv.Atoi(eloStr)
	if err != nil {
		slog.Error("Failed to parse elo", "elo", eloStr, "error", err)
		return -1
	}

	return elo
}

// Parses a filter string to a min and optional max elo.
// Examples:
//	"1500" -> min=1500, max=9999
//	"1500-2000" -> min=1500, max=2000
//	"" -> min=0, max=9999
func ParseEloFilter(filter string) (int, int, error) {
	if len(filter) == 0 {
		return 0, 9999, nil
	}

	formatError := fmt.Errorf("Invalid elo filter format %q. Use <min> or <min>-<max>. Example: 1500 or 1500-2000", filter)

	parts := strings.Split(filter, "-")
	if len(parts) == 1 {
		// Only filter for at least min
		min, err := strconv.Atoi(parts[0])
		if err != nil {
			return -1, -1, formatError
		}

		return min, 9999, nil
	} else if len(parts) == 2 {
		// Filter for between min and max (inc)
		min, errMin := strconv.Atoi(parts[0])
		max, errMax := strconv.Atoi(parts[1])
		if errMin != nil || errMax != nil {
			return -1, -1, formatError
		}

		return min, max, nil
	}

	return -1, -1, formatError
}

// Returns true if both players of a provided game are within the given elo range (inclusive).
func FilterElo(game Game, min int, max int) bool {
	return game.WhiteElo >= min && game.WhiteElo <= max &&
		game.BlackElo >= min && game.BlackElo <= max
}
