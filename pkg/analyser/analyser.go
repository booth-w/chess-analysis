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

type PrintOptions struct {
	// If > 0, only print the top N values. Otherwise, print all.
	Top int

	// If true, any values cut off by [PrintOptions.Top], will be added to "other".
	Other bool

	// If true, sort ascending.
	Asc bool

	// If true, print the percentage on the same line.
	PrintPercent bool

	// If true, print the total at the end.
	PrintTotal bool
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func PrintTotalWinsByColour(gamesData parser.GamesData, options PrintOptions) {
	slog.Info("Getting total wins per colour", "options", options)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	labels := []string{"White", "Black", "Draw", "Invalid"}

	// Print table header
	if options.PrintPercent {
		fmt.Fprintln(w, "Colour\tWins\tPercent")
	} else {
		fmt.Fprintln(w, "Colour\tWins")
	}

	for i, wins := range gamesData.Wins {
		if options.PrintPercent {
			percent := float64(wins) / float64(gamesData.TotalGames) * 100
			fmt.Fprintf(w, "%s\t%d\t%.2f%%\n", labels[i], wins, percent)
		} else {
			fmt.Fprintf(w, "%s\t%d\n", labels[i], wins)
		}
	}

	w.Flush()
	if options.PrintTotal {
		fmt.Printf("Total: %d\n", gamesData.TotalGames)
	}
}

var totalValue int64

// Prints a map of orderable interfaces sorted by value.
// If two values are equal, sort by keys lexicographically.
func PrintSortedMap[K cmp.Ordered, V Number](m map[K]V, options PrintOptions) {
	slog.Info("Printing sorted map", "options", options)

	type kv struct {
		Key   K
		Value V
	}

	var sorted []kv
	for k, v := range m {
		sorted = append(sorted, kv{k, v})
	}

	if options.PrintTotal || options.PrintPercent {
		totalValue = 0
		for _, kv := range sorted {
			totalValue += int64(kv.Value)
		}
	}

	// Sort by value asc/dec, then by key ascending
	// If sorting gets too slow, only sort the top N values
	// For now, sort the whole map
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Value == sorted[j].Value {
			return sorted[i].Key < sorted[j].Key
		}
		return sorted[i].Value > sorted[j].Value
	})

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Print table header
	if options.PrintPercent {
		fmt.Fprintln(w, "Key\tValue\tPercent")
	} else {
		fmt.Fprintln(w, "Key\tValue")
	}

	for _, kv := range sorted {
		if options.PrintPercent {
			percent := float64(kv.Value) / float64(totalValue) * 100
			fmt.Fprintf(w, "%v\t%v\t%.2f%%\n", kv.Key, kv.Value, percent)
		} else {
			fmt.Fprintf(w, "%v\t%v\n", kv.Key, kv.Value)
		}
	}

	w.Flush()
	if options.PrintTotal {
		fmt.Printf("Total: %d\n", totalValue)
	}
}
