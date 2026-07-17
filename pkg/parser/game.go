package parser

type Game struct {
	Movetext []string

	Event string
	Site  string
	Date  string
	Round string
	White string
	Black string

	// 0: white win, 1: black win, 2: draw
	Result int

	UTCDate string
	UTCTime string

	WhiteElo        int
	BlackElo        int
	WhiteRatingDiff int
	BlackRatingDiff int
	WhiteTitle      string
	BlackTitle      string

	ECO         string
	Opening     string
	TimeControl string
	Termination string
	LichessId   string
}

type GamesData struct {
	TotalGames int
	Wins       [3]int

	TimeControls map[string]int
	Terminations map[string]int
}
