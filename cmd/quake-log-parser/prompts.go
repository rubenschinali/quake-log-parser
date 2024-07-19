package main

import (
	"bufio"
	"fmt"
	"os"
	"quake-log-parser/internal/parser"
	"quake-log-parser/internal/report"
	"quake-log-parser/logger"
)

func logPrompt() ([]string, error) {
	// Prompt for log files
	fmt.Print("Enter the log files (space-separated) or directory pattern (e.g., /logs/*.log): ")

	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		logFilesInput := scanner.Text()

		// Expand the input into a list of log files
		logFilesSlice, err := parser.ExpandLogFiles(logFilesInput)
		if err != nil {
			logger.ErrorLogger.Println("Error expanding log files:", err)
			return nil, fmt.Errorf("Error expanding log files: %w", err)
		}
		return logFilesSlice, nil
	}
	// Handle errors from the scanner
	if err := scanner.Err(); err != nil {
		logger.ErrorLogger.Println("Failed to read log files input:", err)
		return nil, fmt.Errorf("failed to read log files input: %w", err)
	}

	logger.ErrorLogger.Println("Failed to read log files input.")
	return nil, fmt.Errorf("failed to read log files input")
}

// outputPrompt prompts the user for the output type (json or screen) and returns the selected type.
func outputPrompt() (string, error) {
	var outputType string
	validOutputs := map[string]bool{
		"json":   true,
		"screen": true,
	}

	// Prompt for output type until a valid option is provided
	for {
		fmt.Print("Enter the output type (json or screen): ")
		fmt.Scanln(&outputType)

		if _, ok := validOutputs[outputType]; ok {
			return outputType, nil
		}
		fmt.Println("Invalid output type. Please enter 'json' or 'screen'.")
	}

}

func handleOutput(outputType string, allGames map[int]*parser.Game) {
	switch outputType {
	case "json":
		outputJson(allGames)
	case "screen":
		outputScreen(allGames)
	default:
		fmt.Println("Invalid output type. Use 'json' or 'screen'.")
		os.Exit(1)
	}
}

func outputJson(allGames map[int]*parser.Game) {
	var jsonFile string
	fmt.Print("Enter the path to save the JSON file: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		jsonFile = scanner.Text()
	} else {
		err := scanner.Err()
		if err != nil {
			logger.ErrorLogger.Println("Error reading input:", err)
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}
	}

	// Validate the file path (optional but recommended)
	if jsonFile == "" {
		fmt.Println("Invalid file path.")
		os.Exit(1)
	}

	err := report.WriteGameReportJSON(jsonFile, allGames)
	if err != nil {
		logger.ErrorLogger.Println("Error writing JSON file:", err)
		fmt.Println("Error writing JSON file:", err)
		os.Exit(1)
	}

	logger.InfoLogger.Println("Game report JSON file written successfully.")
	fmt.Println("Game report JSON file written successfully.")
}

func outputScreen(allGames map[int]*parser.Game) {
	report.ReportScreenOutput(allGames)
}
