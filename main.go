package main

import (
	"flag"
	"log/slog"
	"os"
	"runtime/pprof"

	"github.com/booth-w/chess-analysis/pkg/analyser"
	"github.com/booth-w/chess-analysis/pkg/parser"
)

func main() {
	flagElo := flag.String("elo", "", "Elo rating to filter games by. Usage: <min> or <min>-<max> (inclusive). Example: 1500 or 1500-2000")
	flagProfile := flag.Bool("profile", false, "Enable CPU profiling (creates cpu.prof)")
	flag.Parse()

	if *flagProfile {
		profFile, _ := os.Create("cpu.prof")
		pprof.StartCPUProfile(profFile)
		defer pprof.StopCPUProfile()
	}

	eloMin, eloMax, err := parser.ParseEloFilter(*flagElo)
	if err != nil {
		slog.Error("Error parsing elo filter", "error", err)
		os.Exit(1)
	}

	games := parser.ParseStdin(eloMin, eloMax)

	options := analyser.PrintOptions{
		PrintTotal: true,
	}
	analyser.PrintTotalWinsByColour(games, options)
	analyser.PrintSortedMap(games.Terminations, options)

	slog.Info("Done")
}
