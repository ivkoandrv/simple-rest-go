package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadConfig() error {
	env := os.Getenv("APP_ENV")

	err := godotenv.Load(fmt.Sprintf(".env.%s", env))
	if err != nil {
		return err
	}

	return nil
}
