package analyser

import (
	"fmt"
	"log/slog"
	"os"
	"text/tabwriter"

	"github.com/booth-w/chess-analysis/pkg/parser"
)

func PrintTotalWinsByColour(gamesData parser.GamesData) {
	slog.Info("Getting total wins per colour")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	labels := []string{"White", "Black", "Draw"}

	fmt.Fprintln(w, "Colour\tWins\tPercent")
	for i, wins := range gamesData.Wins {
		percent := float64(wins) / float64(gamesData.TotalGames) * 100
		fmt.Fprintf(w, "%s\t%d\t%.2f%%\n", labels[i], wins, percent)
	}

	w.Flush()
	fmt.Printf("\nTotal: %d\n", gamesData.TotalGames)
}
