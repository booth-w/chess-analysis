package parser

func getWinner(line string) int {
	// Expected inputs:
	// 1-0
	// 0-1
	// 1/2-1/2
	line = parseGeneric(line)

	if len(line) == 7 {
		return 2 // draw
	} else if line[0] == '1' {
		return 0 // white
	} else {
		return 1 // black
	}
}
