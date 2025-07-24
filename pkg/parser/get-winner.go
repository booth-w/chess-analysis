package parser

func GetWinner(line string) (int) {
	if len(line) >= 2 && line[1] == 'R' {
		if line[10] == '/' {
			return 2 // draw
		} else if line[9] == '1' {
			return 0 // white
		} else {
			return 1 // black
		}
	}
	return -1
}
