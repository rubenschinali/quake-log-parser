package parser

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"quake-log-parser/logger"
	"regexp"
	"strings"
	"sync"
)

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

// ParseLogFiles concurrently parses multiple log files and returns a map of parsed games and any encountered errors.
func ParseLogFiles(logFilesSlice []string) (map[int]*Game, error) {
	// Concurrently parse log files
	var wg sync.WaitGroup
	gamesChannel := make(chan map[int]*Game, len(logFilesSlice))
	var errors []error

	for _, logFile := range logFilesSlice {
		wg.Add(1)
		go func(logFile string) {
			defer wg.Done()
			games, err := parseLogFile(logFile)
			if err != nil {
				logger.ErrorLogger.Println("Error parsing log file:", err)
				errors = append(errors, err)
				return
			}
			gamesChannel <- games
		}(logFile)
	}

	wg.Wait()
	close(gamesChannel)

	// Aggregate all games from the channel
	allGames := make(map[int]*Game)
	counter := 0
	for games := range gamesChannel {
		for _, game := range games {
			allGames[counter] = game
			counter++
		}
	}

	// Check if there were any errors during parsing
	var err error
	if len(errors) > 0 {
		err = fmt.Errorf("encountered %d errors during parsing", len(errors))
	}

	return allGames, err
}

// expandLogFiles expands a space-separated list of log files or a directory pattern into a slice of file paths.
func ExpandLogFiles(input string) ([]string, error) {
	if strings.Contains(input, "*") {
		return filepath.Glob(input)
	}
	return strings.Fields(input), nil
}
