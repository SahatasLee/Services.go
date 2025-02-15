package config

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("./dev.env")

	if err != nil {
		slog.Warn("Error loading .env file")
	}
}
