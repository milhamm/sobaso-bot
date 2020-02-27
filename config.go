package main

import "os"

type LineBotConfig struct{
	CHANNEL_SECRET string
	CHANNEL_TOKEN string
}

type Config struct{
	Line LineBotConfig
}

func NewConfig() *Config{
	return &Config{
		Line :LineBotConfig{
			CHANNEL_SECRET: getEnv("CHANNEL_SECRET"),
			CHANNEL_TOKEN: getEnv("CHANNEL_TOKEN"),
		},
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	
	return ""
}