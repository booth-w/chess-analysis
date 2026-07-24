package main

import (
	"encoding/gob"
	"flag"
	"log/slog"
	"os"
	"runtime/pprof"

	"github.com/booth-w/chess-analysis/pkg/analyser"
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

func saveToGob(games parser.GamesData, filepath string) error {
	slog.Info("Saving to gob", "filepath", filepath)

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(games)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	flagElo := flag.String("elo", "", "Elo rating to filter games by. Usage: <min> or <min>-<max> (inclusive). Example: 1500 or 1500-2000")
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

	games, err := parser.ParseStdin(eloMin, eloMax)

	if err != nil {
		slog.Error("Error parsing stdin", "error", err)
		os.Exit(1)
	}

	options := analyser.PrintOptions{
		PrintTotal:   true,
		PrintPercent: true,
	}
	analyser.PrintTotalWinsByColour(games, options)
	analyser.PrintSortedMap(games.Terminations, options)

	// Save to gob
	if *flagGobOut != "" {
		err = saveToGob(games, *flagGobOut)
		if err != nil {
			slog.Error("Error saving to gob", "error", err)
			os.Exit(1)
		}
	}

	slog.Info("Done")
}
