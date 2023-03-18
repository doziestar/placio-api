package config

import (
	"context"
	"github.com/spf13/viper"
	"placio-pkg/logger"
	"time"
)

type Config struct {
	JwtSecret    string        `map structure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `map structure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `map structure:"JWT_MAX_AGE"`
	DatabaseURL  string        `map structure:"DATABASE_URL"`
	ClientOrigin string        `map structure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	logger.Info(context.Background(), "Loading config from: "+path)
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	// print the full path of the config file
	logger.Info(context.Background(), "Config file path: "+viper.ConfigFileUsed())

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
