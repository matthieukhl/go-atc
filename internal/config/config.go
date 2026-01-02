package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenSkyApiKey   string
	OpenSkyUsername string
	OpenSkyPassword string
}

func NewConfig() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, err
	}
	return Config{
			OpenSkyApiKey:   os.Getenv("OPENSKY_API_TOKEN"),
			OpenSkyUsername: os.Getenv("OPENSKY_USERNAME"),
			OpenSkyPassword: os.Getenv("OPENSKY_PASSWORD"),
		},
		nil
}
