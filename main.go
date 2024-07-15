package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

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

// parseLogFile tries to open a file and then parses each line to a predefined regex to match Players, Kills and KillMethods
// It returns a map of game IDs to Game objects and any error encountered during file parsing.
func parseLogFile(filename string) (map[int]*Game, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := make(map[int]*Game)
	var currentGame *Game
	currentGameID := 0

	// killRegex matches lines formatted as:
	//   "  Kill: <timestamp>: <attacker> killed <victim> by <killMethod>"
	killRegex := regexp.MustCompile(
		`\s+Kill:\s+\d+\s+\d+\s+\d+:\s+` + // "Kill": <timestamp>:
			`([^ ]+)` + // <attacker>
			`\s+killed\s+` + // "killed"
			`([^ ]+)` + // <victim>
			`\s+by\s+` + // "by"
			`([^\s]+)`) // <killMethod>

	// Scanning throughout the file
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "InitGame") {
			currentGameID++
			// Initializing NewGame
			currentGame = NewGame()
			games[currentGameID] = currentGame

			// Parsing "Kill" line
		} else if strings.Contains(line, "Kill:") {
			match := killRegex.FindStringSubmatch(line)
			if match != nil && currentGame != nil {
				attacker := match[1]
				victim := match[2]
				killMethod := match[3]

				if attacker != "<world>" {
					currentGame.Players[attacker] = struct{}{}
					currentGame.Kills[attacker]++
				} else {
					// Subtracts one score point if Player was killed by <world>
					currentGame.Kills[victim]--
				}

				currentGame.Players[victim] = struct{}{}
				currentGame.TotalKills++
				currentGame.KillMethods[killMethod]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return games, nil
}

// writeGameReportJSON writes the games map to a json file, where each line is a Game read from the .log file.
func writeGameReportJSON(games map[int]*Game, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	for _, game := range games {
		gameData := map[string]interface{}{
			"total_kills":    game.TotalKills,
			"players":        game.Players,
			"kills":          game.Kills,
			"kills_by_means": game.KillMethods,
		}

		err := encoder.Encode(gameData)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	games, err := parseLogFile("qgames.log")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = writeGameReportJSON(games, "game_report.json")
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("Game report JSON file written successfully.")
}
