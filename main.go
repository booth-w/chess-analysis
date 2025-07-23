package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	total := 0
	wins := [3]int{}
	for s.Scan() {
		line := s.Text()

		if len(line) >= 2 && line[1] == 'R' {
			total++
			if line[10] == '/' {
				wins[2]++
			} else if line[9] == '1' {
				wins[0]++
			} else {
				wins[1]++
			}
		}
	}
	fmt.Printf("White: %d, Black: %d, Draw: %d, Total: %d\n", wins[0], wins[1], wins[2], total)
}
