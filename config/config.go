package config

import (
	"log"
	"os"
	d "github.com/milhamm/sobaso-bot/domain"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)



func NewBot(channelSecret, channelToken string) (*d.LineBot, error){
	bot, err := linebot.New(
		channelSecret,
		channelToken,
	)

	if err != nil {
		return nil, err
	}

	return &d.LineBot{
		bot	:	bot,
	}, nil
}

func New() *d.LineBotConfig{

	if err := godotenv.Load(); err != nil {
        log.Print(err)
    }

	return &d.LineBotConfig{
		CHANNEL_SECRET_CLIENT: getEnv("CHANNEL_SECRET_CLIENT"),
		CHANNEL_TOKEN_CLIENT: getEnv("CHANNEL_TOKEN_CLIENT"),
		CHANNEL_SECRET_ADMIN: getEnv("CHANNEL_SECRET_ADMIN"),
		CHANNEL_TOKEN_ADMIN: getEnv("CHANNEL_TOKEN_ADMIN"),
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}