package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/booth-w/chess-analysis/pkg/parser"
)

func main() {
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
