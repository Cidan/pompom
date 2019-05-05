package main

import (
	"os"

	"github.com/Cidan/pompom/settings"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	setupLogging()
	settings.Setup("")
}

func setupLogging() {
	var level zerolog.Level
	switch viper.GetString("level") {
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	case "debug":
		level = zerolog.DebugLevel
	default:
		level = zerolog.InfoLevel
	}
	// If we're in a terminal, pretty print
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).Level(level)
		log.Info().Msg("Detected terminal, pretty logging enabled.")
	}
}
