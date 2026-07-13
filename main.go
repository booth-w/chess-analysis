package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/pprof"

	"github.com/booth-w/chess-analysis/pkg/parser"
)

func main() {
	flagProfile := flag.Bool("profile", false, "Enable CPU profiling (creates cpu.prof)")
	flag.Parse()

	if *flagProfile {
		profFile, _ := os.Create("cpu.prof")
		pprof.StartCPUProfile(profFile)
		defer pprof.StopCPUProfile()
	}

	games := parser.ParseStdin()

	slog.Info("Getting total wins per colour")

	whitePercent := float64(games.Wins[0]) / float64(games.TotalGames) * 100
	blackPercent := float64(games.Wins[1]) / float64(games.TotalGames) * 100
	drawPercent := float64(games.Wins[2]) / float64(games.TotalGames) * 100

	fmt.Printf(
		"White: %d (%.2f%%)\nBlack: %d (%.2f%%)\nDraw:  %d (%.2f%%)\nTotal: %d\n",
		games.Wins[0], whitePercent,
		games.Wins[1], blackPercent,
		games.Wins[2], drawPercent,
		games.TotalGames,
	)

	slog.Info("Done")
}
