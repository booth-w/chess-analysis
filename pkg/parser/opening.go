package parser

import (
	"fmt"
	"strings"

	"github.com/booth-w/chess-analysis/pkg/game"
)

func parseOpening(line string) (game.Opening, error) {
	line, err := parseGeneric(line)
	if err != nil {
		return game.Opening{}, err
	}

	openingParts := strings.Split(line, ": ")
	if len(openingParts) > 2 {
		return game.Opening{}, fmt.Errorf("invalid opening format %q. Expected <family> or <family>:<variation>", line)
	}

	opening := game.Opening{
		Family: openingParts[0],
	}

	if len(openingParts) == 2 {
		variation := strings.Split(openingParts[1], ", ")
		opening.Variation = variation
	}

	return opening, nil
}
