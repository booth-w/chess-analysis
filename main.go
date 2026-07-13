package main

import (
	"bufio"
	"flag"
	"fmt"
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

	s := bufio.NewScanner(os.Stdin)

	game := [2][]string{}
	totalGames := 0
	wins := [3]int{}

	// itterate over each line
	for s.Scan() {
		line := s.Text()
		if len(line) != 0 {
			if line[0] == '[' {
				game[0] = append(game[0], line)
			} else {
				game[1] = append(game[1], line)
			}
			continue
		} else if len(game[1]) == 0 {
			continue
		}

		winner := parser.GetWinner(game[0])
		if winner != -1 {
			totalGames++
			wins[winner]++
		}
		game = [2][]string{}
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
}
