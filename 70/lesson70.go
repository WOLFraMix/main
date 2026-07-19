package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	log.Debug().Msg("Это отладочное сообщение")
	log.Info().Msg("Это информационное сообщение")
	log.Warn().Msg("Это предупреждение")
	log.Error().Msg("Это сообщение об ошибке")
	log.Fatal().Msg("Это фатальная ошибка")
}
