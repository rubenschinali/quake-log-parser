package main

import (
	"fmt"
	"os"

	"quake-log-parser/internal/parser"
	"quake-log-parser/logger"
)

func main() {
	logger.Init()

	startCLI()

	logFilesSlice, err := logPrompt()
	if err != nil {
		logger.ErrorLogger.Println("Error getting log files:", err)
		os.Exit(1)
	}

	// Get the output type from the user
	outputType, err := outputPrompt()
	if err != nil {
		logger.ErrorLogger.Println("Error getting output type:", err)
		os.Exit(1)
	}

	allGames, err := parser.ParseLogFiles(logFilesSlice)
	if err != nil {
		logger.ErrorLogger.Println("Error parsing log files:", err)
		os.Exit(1)
	}

	handleOutput(outputType, allGames)

	fmt.Println("THE END!")
}
