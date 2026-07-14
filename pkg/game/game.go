package game

type Game struct {
	Movetext []string

	Event  string
	Site   string
	Date   string
	Round  string
	White  string
	Black  string
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
