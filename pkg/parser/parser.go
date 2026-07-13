package parser

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/booth-w/chess-analysis/pkg/game"
)

type GamesData struct {
	TotalGames int
	Wins       [3]int
}

func ParseStdin() GamesData {
	slog.Info("Reading stdin")
	s := bufio.NewScanner(os.Stdin)

	var gamesData GamesData
	var newGame game.Game

	for s.Scan() {
		line := s.Text()

		if len(line) != 0 {
			if line[0] == '[' {
				metadataKey := strings.Split(line, " ")[0][1:]

				switch metadataKey {
				case "Event":
					newGame.Event = parseGeneric(line)
				case "Site":
					newGame.Site = parseGeneric(line)
				case "Date":
					newGame.Date = parseGeneric(line)
				case "Round":
					newGame.Round = parseGeneric(line)
				case "White":
					newGame.White = parseGeneric(line)
				case "Black":
					newGame.Black = parseGeneric(line)
				case "Result":
					newGame.Result = getWinner(parseGeneric(line))
				case "UTCDate":
					newGame.UTCDate = parseGeneric(line)
				case "UTCTime":
					newGame.UTCTime = parseGeneric(line)
				case "WhiteElo":
					// newGame.WhiteElo = parseGenericInt(line)
				case "BlackElo":
					// newGame.BlackElo = parseGenericInt(line)
				case "WhiteRatingDiff":
					newGame.WhiteRatingDiff = parseGenericInt(line)
				case "BlackRatingDiff":
					newGame.BlackRatingDiff = parseGenericInt(line)
				case "WhiteTitle":
					newGame.WhiteTitle = parseGeneric(line)
				case "BlackTitle":
					newGame.BlackTitle = parseGeneric(line)
				case "ECO":
					newGame.ECO = parseGeneric(line)
				case "Opening":
					newGame.Opening = parseGeneric(line)
				case "TimeControl":
					newGame.TimeControl = parseGeneric(line)
				case "Termination":
					newGame.Termination = parseGeneric(line)
				case "LichessId":
					newGame.LichessId = parseGeneric(line)
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

		gamesData.TotalGames++
		gamesData.Wins[newGame.Result]++
		newGame = game.Game{}
	}

	return gamesData
}

func parseGeneric(line string) string {
	return strings.Split(line, "\"")[1]
}

func parseGenericInt(line string) int {
	out, err := strconv.Atoi(parseGeneric(line))
	if err != nil {
		slog.Error("Error parsing PGN to int", "line", line)
	}

	return out
}
