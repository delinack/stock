package config

import (
	"github.com/spf13/viper"
	"storage/internal/pkg/logger"
	"storage/internal/pkg/storage"
)

// ApplicationConfig the main config for app
type ApplicationConfig struct {
	HTTPServerPort      string `mapstructure:"APP_PORT" validate:"required"`
	storage.PGConfig    `mapstructure:",squash"`
	logger.LoggerConfig `mapstructure:",squash"`
}

// MustParseConfig is main config parser
func MustParseConfig() ApplicationConfig {
	viper.SetDefault("POSTGRES_DB", "stock")
	viper.SetDefault("POSTGRES_HOST", "localhost")
	viper.SetDefault("LOG_LEVEL", "warn")

	cfg := ApplicationConfig{}
	if err := ParseConfig(&cfg); err != nil {
		panic(err)
	}

	return cfg
}
