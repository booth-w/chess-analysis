package parser

// Expected inputs:
//	[Result "1-0"]
//	[Result "0-1"]
//	[Result "1/2-1/2"]
//
// Returns:
//	0: white win
//	1: black win
//	2: draw
func getWinner(line string) int {
	line = parseGeneric(line)

	if len(line) == 7 {
		return 2 // draw
	} else if line[0] == '1' {
		return 0 // white
	} else {
		return 1 // black
	}
}
