package parser

import (
	"testing"
)

func TestParseLogFile(t *testing.T) {
	games, err := ParseLogFile("testdata/qgames.log")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(games) == 0 {
		t.Errorf("Expected some games, got %v", len(games))
	}

	for id, game := range games {
		t.Logf("Game ID: %d, Total Kills: %d", id, game.TotalKills)
	}
}

func TestWriteGameReportJSON(t *testing.T) {
	games, err := ParseLogFile("testdata/qgames.log")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = WriteGameReportJSON(games, "testdata/game_report.json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
