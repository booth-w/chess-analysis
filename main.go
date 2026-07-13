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
	totalGames := len(games)
	wins := [3]int{}

	for _, g := range games {
		wins[parser.GetWinner(g.Result)]++
	}

	whitePercent := float64(wins[0]) / float64(totalGames) * 100
	blackPercent := float64(wins[1]) / float64(totalGames) * 100
	drawPercent := float64(wins[2]) / float64(totalGames) * 100

	fmt.Printf(
		"White: %d (%.2f%%)\nBlack: %d (%.2f%%)\nDraw:  %d (%.2f%%)\nTotal: %d\n",
		wins[0], whitePercent,
		wins[1], blackPercent,
		wins[2], drawPercent,
		totalGames,
	)

	slog.Info("Done")
}
