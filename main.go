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

	fmt.Println(*flagProfile)

	if *flagProfile {
		profFile, _ := os.Create("cpu.prof")
		pprof.StartCPUProfile(profFile)
		defer pprof.StopCPUProfile()
	}

	s := bufio.NewScanner(os.Stdin)

	total := 0
	wins := [3]int{}
	for s.Scan() {
		line := s.Text()

		winner := parser.GetWinner(line)
		if (winner != -1) {
			total++
			wins[winner]++
		}
	}
	fmt.Printf("White: %d\nBlack: %d\nDraw:  %d\nTotal: %d\n", wins[0], wins[1], wins[2], total)
}
