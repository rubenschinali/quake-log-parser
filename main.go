package main

import (
	"flag"
	"fmt"
	"os"

	"quake-log-parser/internal/parser"
	"quake-log-parser/log"
)

func main() {
	log.Init()

	logFile := flag.String("logfile", "qgames.log", "Path to the log file")
	jsonFile := flag.String("jsonfile", "game_report.json", "Path to the output JSON file")
	flag.Parse()

	games, err := parser.ParseLogFile(*logFile)
	if err != nil {
		log.ErrorLogger.Println("Error parsing log file:", err)
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	err = parser.WriteGameReportJSON(games, *jsonFile)
	if err != nil {
		log.ErrorLogger.Println("Error writing JSON file:", err)
		fmt.Println("Error writing JSON file:", err)
		os.Exit(1)
	}

	log.InfoLogger.Println("Game report JSON file written successfully.")
	fmt.Println("Game report JSON file written successfully.")
}
