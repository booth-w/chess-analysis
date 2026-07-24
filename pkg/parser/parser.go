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
				metadataKey := line[1:strings.Index(line, " ")]
				err := error(nil)

				switch metadataKey {
				case "Event":
					newGame.Event, err = parseGeneric(line)
				case "Site":
					newGame.Site, err = parseGeneric(line)
				case "Date":
					newGame.Date, err = parseGeneric(line)
				case "Round":
					newGame.Round, err = parseGeneric(line)
				case "White":
					newGame.White, err = parseGeneric(line)
				case "Black":
					newGame.Black, err = parseGeneric(line)
				case "Result":
					newGame.Result, err = getWinner(line)
				case "UTCDate":
					newGame.UTCDate, err = parseGeneric(line)
				case "UTCTime":
					newGame.UTCTime, err = parseGeneric(line)
				case "WhiteElo":
					newGame.WhiteElo, err = parseElo(line)
				case "BlackElo":
					newGame.BlackElo, err = parseElo(line)
				case "WhiteRatingDiff":
					newGame.WhiteRatingDiff, err = parseGenericInt(line)
				case "BlackRatingDiff":
					newGame.BlackRatingDiff, err = parseGenericInt(line)
				case "WhiteTitle":
					newGame.WhiteTitle, err = parseGeneric(line)
				case "BlackTitle":
					newGame.BlackTitle, err = parseGeneric(line)
				case "ECO":
					newGame.ECO, err = parseGeneric(line)
				case "Opening":
					newGame.Opening, err = parseGeneric(line)
				case "TimeControl":
					newGame.TimeControl, err = parseGeneric(line)
				case "Termination":
					newGame.Termination, err = parseGeneric(line)
				case "LichessId":
					newGame.LichessId, err = parseGeneric(line)
				default:
					slog.Warn("Unknown metadata key", "metadataKey", metadataKey)
				}

				if err != nil {
					slog.Error("Failed to parse metadata", "line", line, "error", err, "game", newGame)
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

		if newGame.Result != -1 {
			gamesData.TotalGames++
			gamesData.Wins[newGame.Result]++
		}

		newGame = Game{}
	}

	return gamesData
}

// Parses a PGN metadata line and returns the value between quotes.
func parseGeneric(line string) (string, error) {
	start := strings.IndexByte(line, '"')
	if start == -1 {
		return "", fmt.Errorf("invalid PGN line: %s", line)
	}

	end := strings.IndexByte(line[start+1:], '"')
	if end == -1 {
		return "", fmt.Errorf("invalid PGN line: %s", line)
	}

	return line[start+1 : start+1+end], nil
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
