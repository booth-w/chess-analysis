package analyser

import (
	"fmt"
	"log/slog"

	"github.com/booth-w/chess-analysis/pkg/parser"
)

func PrintTotalWinsByColour(gamesData parser.GamesData) {
	slog.Info("Getting total wins per colour")

	whitePercent := float64(gamesData.Wins[0]) / float64(gamesData.TotalGames) * 100
	blackPercent := float64(gamesData.Wins[1]) / float64(gamesData.TotalGames) * 100
	drawPercent := float64(gamesData.Wins[2]) / float64(gamesData.TotalGames) * 100

	fmt.Printf(
		"White: %d (%.2f%%)\nBlack: %d (%.2f%%)\nDraw:  %d (%.2f%%)\nTotal: %d\n",
		gamesData.Wins[0], whitePercent,
		gamesData.Wins[1], blackPercent,
		gamesData.Wins[2], drawPercent,
		gamesData.TotalGames,
	)
}
