package gob

import (
	"encoding/gob"
	"log/slog"
	"os"

	"github.com/booth-w/chess-analysis/pkg/parser"
)

func SaveToGob(games parser.GamesData, filepath string) error {
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

func LoadFromGob(filepath string) (parser.GamesData, error) {
	slog.Info("Loading from gob", "filepath", filepath)

	file, err := os.Open(filepath)
	if err != nil {
		return parser.GamesData{}, err
	}
	defer file.Close()

	var games parser.GamesData
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&games)
	if err != nil {
		return parser.GamesData{}, err
	}

	return games, nil
}
