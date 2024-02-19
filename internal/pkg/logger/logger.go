package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// LoggerConfig struct for logger config
type LoggerConfig struct {
	Level string `mapstructure:"LOG_LEVEL"`
}

// NewLogger constructor for logger
func NewLogger(cfg LoggerConfig) zerolog.Logger {
	w := zerolog.NewConsoleWriter()
	w.FormatTimestamp = func(i interface{}) string {
		return time.Now().UTC().Format(time.Stamp)
	}

	level, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		panic(fmt.Errorf("zerolog.ParseLevel failed: %w", err))
	}

	w.Out = os.Stderr
	logger := zerolog.New(w).Level(level).With().Timestamp().Logger()

	return logger
}
