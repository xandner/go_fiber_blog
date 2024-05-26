package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger interface {
	Info(message string, fields ...interface{})
	Error(message string, err error)
	Run()
}

type zeroLogger struct {
	logger *zerolog.Logger
}

func (l *zeroLogger) Info(message string, fields ...interface{}) {
	l.logger.Info().Msgf(message, fields...)
}

func (l *zeroLogger) Error(message string, err error) {
	l.logger.Error().Err(err).Msg(message)
}
func (l *zeroLogger) Run() {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	duration := next.Sub(now)
	time.AfterFunc(duration, func() {
		l.Run()
	})
}
func New(level string) (Logger, *zerolog.Logger) {
	logWriter := zerolog.MultiLevelWriter(os.Stdout)
	logger := zerolog.New(logWriter).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount).Logger()
	return &zeroLogger{logger: &logger}, &logger
}
