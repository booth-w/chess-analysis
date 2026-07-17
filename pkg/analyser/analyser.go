package analyser

import (
	"cmp"
	"fmt"
	"log/slog"
	"os"
	"sort"
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

// Prints a map of string int sorted desc by value.
// If two values are equal, sort by keys lexicographically.
func PrintSortedMap[K cmp.Ordered, V cmp.Ordered](m map[K]V) {
	slog.Info("Printing sorted map")

	type kv struct {
		Key   K
		Value V
	}

	var sorted []kv
	for k, v := range m {
		sorted = append(sorted, kv{k, v})
	}

	// Sort by value descending, then by key ascending
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Value == sorted[j].Value {
			return sorted[i].Key < sorted[j].Key
		}
		return sorted[i].Value > sorted[j].Value
	})

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, kv := range sorted {
		fmt.Fprintf(w, "%v\t%v\n", kv.Key, kv.Value)
	}
	w.Flush()
}
