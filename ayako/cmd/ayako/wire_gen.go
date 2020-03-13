// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/deissh/osu-lazer/ayako/app"
	"github.com/deissh/osu-lazer/ayako/config"
	"github.com/deissh/osu-lazer/ayako/store/sql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// Injectors from main.go:

func Injector(configPath string) *app.App {
	configConfig := config.Init(configPath)
	store := sql.Init(configConfig)
	appApp := app.NewApp(configConfig, store)
	return appApp
}

// main.go:

var Version string

var Commit string

var Branch string

var BuildTimestamp string

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:     os.Stderr,
		NoColor: false,
	},
	).With().Caller().Logger()
	log.Info().
		Str("version", Version).
		Str("branch", Branch).
		Str("commit", Commit).
		Str("build_timestamp", BuildTimestamp).
		Send()
	log.Debug().Msg("Start initialize dependencies")
	app2 := Injector("config.yaml")
	log.Debug().Msg("Initialize dependencies successful done")

	if err := app2.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
