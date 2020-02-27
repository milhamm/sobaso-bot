package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type LineBotConfig struct{
	CHANNEL_SECRET string
	CHANNEL_TOKEN string
}

func New() *LineBotConfig{

	if err := godotenv.Load(); err != nil {
        log.Print(err)
    }

	return &LineBotConfig{
		CHANNEL_SECRET: getEnv("CHANNEL_SECRET"),
		CHANNEL_TOKEN: getEnv("CHANNEL_TOKEN"),
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}