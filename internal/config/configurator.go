package config

import (
	"errors"
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// ParseConfig parse the config into a struct
func ParseConfig(cfg interface{}) error {
	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	viper.AddConfigPath(workingDir)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		// if config cannot be read, use env variables
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return err
		}
	}

	envKeysMap := map[string]interface{}{}
	if err := mapstructure.Decode(cfg, &envKeysMap); err != nil {
		return err
	}

	for key := range envKeysMap {
		if err := viper.BindEnv(key); err != nil {
			return err
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	validate := validator.New()

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("mapstructure"), ",", 2)[0]
		return name
	})

	if err := validate.Struct(cfg); err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return errors.New("invalid value for validation")
		}

		var errs error
		for _, err := range err.(validator.ValidationErrors) {
			errs = errors.Join(errs, err)
		}
		return errs
	}

	return nil
}
