package report

import (
	"fmt"
	"quake-log-parser/internal/parser"
	"strings"
)

// PrintPretty prints the parsed data in a user-friendly format.
func ScreenOutput(allGames map[int]*parser.Game) {
	for id, game := range allGames {
		fmt.Printf("Game ID: %d\n", id)
		fmt.Println("Game Details:")
		fmt.Printf("  Total Kills: %d\n", game.TotalKills)

		fmt.Println("  Players:")
		for player := range game.Players {
			fmt.Printf("    %s\n", player)
		}

		fmt.Println("  Kills:")
		for player, kills := range game.Kills {
			fmt.Printf("    %s: %d kills\n", player, kills)
		}

		fmt.Println("  Kill Methods:")
		for method, count := range game.KillMethods {
			fmt.Printf("    %s: %d kills\n", method, count)
		}

		fmt.Println(strings.Repeat("-", 20))
	}
}
