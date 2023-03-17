package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	UserServiceHost  string
	UserServicePort  string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	Environment      string
	LogLevel         string
}

func Load() Config {
	c := Config{}

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToString(getOrReturnDefault("USER_SERVICE_PORT", ":8083"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "ahrorbek"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "3108"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToString(getOrReturnDefault("POSTGRES_PORT", "5432"))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "user_db"))
	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{}{
	_, exists := os.LookupEnv(key)
	if exists{
		return os.Getenv(key)
	}

	return defaultValue
}
