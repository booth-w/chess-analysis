package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/pprof"

	"github.com/booth-w/chess-analysis/pkg/analyser"
	"github.com/booth-w/chess-analysis/pkg/game"
	"github.com/booth-w/chess-analysis/pkg/gob"
	"github.com/booth-w/chess-analysis/pkg/parser"
)

func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	case "":
		return slog.Level(99)
	default:
		return slog.LevelInfo
	}
}

func main() {
	flagElo := flag.String("elo", "", "Elo rating to filter games by. Usage: <min> or <min>-<max> (inclusive). Example: 1500 or 1500-2000")
	flagGobIn := flag.String("i", "", "Input path for the gob file")
	flagGobOut := flag.String("o", "", "Output path for the gob file")
	flagProfile := flag.Bool("profile", false, "Enable CPU profiling (creates cpu.prof)")
	flagLogLevel := flag.String("log-level", "", "Set the log level. Usage: debug, info, warn, error")
	flag.Parse()

	logLevel := parseLogLevel(*flagLogLevel)

	if *flagLogLevel != "info" && logLevel == slog.LevelInfo {
		slog.Warn("Invalid log level. Defaulting to info", "level", *flagLogLevel)
	}

	slog.SetDefault(slog.New(
		slog.NewTextHandler(
			os.Stderr,
			&slog.HandlerOptions{
				Level: logLevel,
			},
		),
	))

	if *flagProfile {
		profFile, _ := os.Create("cpu.prof")
		pprof.StartCPUProfile(profFile)
		defer pprof.StopCPUProfile()
	}

	eloMin, eloMax, err := parser.ParseEloFilter(*flagElo)
	if err != nil {
		slog.Error("Error parsing elo filter", "error", err)
		os.Exit(1)
	}

	var games game.GamesData
	if *flagGobIn != "" { // Load from gob file
		games, err = gob.LoadFromGob(*flagGobIn)
		if err != nil {
			slog.Error("Error loading from gob", "error", err)
			os.Exit(1)
		}
	} else { // Parse from stdin
		games, err = parser.ParseStdin(eloMin, eloMax)
		if err != nil {
			slog.Error("Error parsing stdin", "error", err)
			os.Exit(1)
		}
	}

	options := analyser.PrintOptions{
		PrintTotal:   true,
		PrintPercent: true,
	}
	fmt.Println("Win rate:")
	analyser.PrintTotalWinsByColour(games, options)
	fmt.Println("\nTerminations:")
	analyser.PrintSortedMap(games.Terminations, options)
	fmt.Println("\nGame Types:")
	analyser.PrintSortedMap(games.Events, options)

	// Save to gob
	if *flagGobOut != "" {
		err = gob.SaveToGob(games, *flagGobOut)
		if err != nil {
			slog.Error("Error saving to gob", "error", err)
			os.Exit(1)
		}
	}

	slog.Info("Done")
}
