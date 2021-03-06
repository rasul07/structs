package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	LogLevel string
	HttpPort string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env fild found", err)
	}
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "rasul"))
	config.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1234"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "roadmap_tasks"))

	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8002"))

	return config

}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
