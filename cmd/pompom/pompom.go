package main

import (
	"context"
	"os"

	"github.com/Cidan/pompom/database"
	"github.com/Cidan/pompom/settings"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	settings.Setup("")
	setupLogging()
	output, err := database.NewPubsub(
		context.Background(),
		viper.GetString("output.pubsub.project"),
		viper.GetString("output.pubsub.topic"),
	)

	// TODO: mux pubsub with a local cache
	if err != nil {
		log.Fatal().Err(err).Msg("unable to start pubsub")
		return
	}

	if viper.GetBool("input.flatfile.enabled") {
		ff, err := database.NewFlatFile(viper.GetString("input.flatfile.location"), true)
		if err != nil {
			log.Fatal().Err(err).Msg("unable to read flat file")
			return
		}
		ff.Start(output)
	}
	log.Info().Msg("pom pom is now running!")
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
