package parser

// Game represents the state of a game session.
type Game struct {
	TotalKills  int                 // TotalKill tracks the overall number of kills in the game.
	Players     map[string]struct{} // Players is a set of unique players in the game.
	Kills       map[string]int      // Kills maps player names to their individual kills.
	KillMethods map[string]int      // KillMethods maps kill methods to the number of kills in each game.
}

// NewGame initializes and returns a new instance of the Game struct.
func NewGame() *Game {
	return &Game{
		Players:     make(map[string]struct{}),
		Kills:       make(map[string]int),
		KillMethods: make(map[string]int),
	}
}

// // GameReport represents the structure of a game report in JSON format.
// type GameReport struct {
// 	TotalKills int            `json:"total_kills"`
// 	Players    []string       `json:"players"`
// 	Kills      map[string]int `json:"kills"`
// }

// Prepare the data for JSON encoding
type GameDetails struct {
	TotalKills  int            `json:"TotalKills"`
	Players     []string       `json:"Players"`
	Kills       map[string]int `json:"Kills"`
	KillMethods map[string]int `json:"KillMethods"`
}

type GameReport struct {
	GameID      string      `json:"GameID"`
	GameDetails GameDetails `json:"GameDetails"`
}
