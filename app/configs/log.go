package configs

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Info(format string, args ...interface{}) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msgf(format, args...)
}

func Debug(format string, args ...interface{}) {
	log.Debug().Msgf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Fatal().Msgf(format, args...)
}

func Warn(format string, args ...interface{}) {
	log.Warn().Msgf(format, args...)
}

func Error(format string, args ...interface{}) {
	log.Error().Msgf(format, args...)
}
