package parser

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func ParseStdin(eloMin int, eloMax int) GamesData {
	slog.Info("Reading stdin")
	s := bufio.NewScanner(os.Stdin)

	var gamesData GamesData
	var newGame Game

	for s.Scan() {
		line := s.Text()

		if len(line) != 0 {
			if line[0] == '[' {
				metadataKey := strings.Split(line, " ")[0][1:]

				switch metadataKey {
				case "Event":
					newGame.Event, _ = parseGeneric(line)
				case "Site":
					newGame.Site, _ = parseGeneric(line)
				case "Date":
					newGame.Date, _ = parseGeneric(line)
				case "Round":
					newGame.Round, _ = parseGeneric(line)
				case "White":
					newGame.White, _ = parseGeneric(line)
				case "Black":
					newGame.Black, _ = parseGeneric(line)
				case "Result":
					newGame.Result, _ = getWinner(line)
				case "UTCDate":
					newGame.UTCDate, _ = parseGeneric(line)
				case "UTCTime":
					newGame.UTCTime, _ = parseGeneric(line)
				case "WhiteElo":
					newGame.WhiteElo, _ = parseElo(line)
				case "BlackElo":
					newGame.BlackElo, _ = parseElo(line)
				case "WhiteRatingDiff":
					newGame.WhiteRatingDiff, _ = parseGenericInt(line)
				case "BlackRatingDiff":
					newGame.BlackRatingDiff, _ = parseGenericInt(line)
				case "WhiteTitle":
					newGame.WhiteTitle, _ = parseGeneric(line)
				case "BlackTitle":
					newGame.BlackTitle, _ = parseGeneric(line)
				case "ECO":
					newGame.ECO, _ = parseGeneric(line)
				case "Opening":
					newGame.Opening, _ = parseGeneric(line)
				case "TimeControl":
					newGame.TimeControl, _ = parseGeneric(line)
				case "Termination":
					newGame.Termination, _ = parseGeneric(line)
				case "LichessId":
					newGame.LichessId, _ = parseGeneric(line)
				default:
					slog.Warn("Unknown metadata key", "metadataKey", metadataKey)
				}
			} else {
				newGame.Movetext = append(newGame.Movetext, line)
			}
			continue
		} else if len(newGame.Movetext) == 0 {
			continue
		}

		// Filter by elo
		if !FilterElo(newGame, eloMin, eloMax) {
			newGame = Game{}
			continue
		}

		// Time control
		if gamesData.TimeControls == nil {
			gamesData.TimeControls = make(map[string]int)
		}
		gamesData.TimeControls[newGame.TimeControl]++

		// Termination
		if gamesData.Terminations == nil {
			gamesData.Terminations = make(map[string]int)
		}
		gamesData.Terminations[newGame.Termination]++

		gamesData.TotalGames++
		gamesData.Wins[newGame.Result]++
		newGame = Game{}
	}

	return gamesData
}

// Parses a PGN metadata line and returns the value between quotes.
func parseGeneric(line string) (string, error) {
	split := strings.Split(line, "\"")

	if len(split) < 2 {
		return "", fmt.Errorf("invalid PGN line: %s", line)
	}

	return split[1], nil
}

// Calls [parseGeneric] and converts to int.
func parseGenericInt(line string) (int, error) {
	generic, err := parseGeneric(line)
	if err != nil {
		return -1, err
	}

	out, err := strconv.Atoi(generic)
	if err != nil {
		return -1, fmt.Errorf("invalid integer in PGN line: %s", line)
	}

	return out, nil
}
