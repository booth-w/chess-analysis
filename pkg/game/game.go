package game

type Game struct {
	Movetext []string

	Event string
	Site  string
	Date  string
	Round string
	White string
	Black string

	// 0: white win, 1: black win, 2: draw, 3: invalid
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
	Opening     Opening
	TimeControl string
	Termination string
	LichessId   string
}

// Example. "Sicilian Defense: Najdorf Variation, English Attack" is represented as:
//
//	{
//		Family: "Sicilian Defense",
//		Variation: ["Najdorf Variation", "English Attack"]
//	}
type Opening struct {
	Family    string
	Variation []string
}

type GamesData struct {
	TotalGames int
	Wins       [4]int

	Events            map[string]int
	Titles            map[string]int
	Openings          map[string]int
	OpeningVariations map[string]int

	TimeControls map[string]int
	Terminations map[string]int
}
