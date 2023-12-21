package configs

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InfoLog(format string, args ...interface{}) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msgf(format, args...)
}
