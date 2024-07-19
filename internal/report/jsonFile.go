package report

import (
	"encoding/json"
	"fmt"
	"os"
	"quake-log-parser/internal/parser"
)

// WriteGameReportJSON saves the parsed data as a JSON file with two-digit game IDs.
func WriteGameReportJSON(filename string, allGames map[int]*parser.Game) error {
	// Prepare the data for JSON encoding
	formattedReport := make(map[string]parser.GameReport)

	// Use an incremental counter for sequential IDs
	counter := 1

	for _, game := range allGames {
		// Convert the Players map to a slice of strings
		players := make([]string, 0, len(game.Players))
		for player := range game.Players {
			players = append(players, player)
		}

		// Create the GameReport entry with two-digit formatted ID
		formattedReport[fmt.Sprintf("game_%02d", counter)] = parser.GameReport{
			TotalKills: game.TotalKills,
			Players:    players,
			Kills:      game.Kills,
		}
		counter++ // Increment the counter for the next game
	}

	// Create and open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating JSON file: %w", err)
	}
	defer file.Close()

	// Encode the data to JSON with pretty-printing
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print
	if err := encoder.Encode(formattedReport); err != nil {
		return fmt.Errorf("error encoding JSON data: %w", err)
	}

	return nil
}
