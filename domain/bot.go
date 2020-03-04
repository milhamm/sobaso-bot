package domain

import "github.com/line/line-bot-sdk-go/linebot"

type LineBot struct {
	bot *linebot.Client
}

type LineBotConfig struct{
	CHANNEL_SECRET_CLIENT string
	CHANNEL_TOKEN_CLIENT string
	CHANNEL_SECRET_ADMIN string
	CHANNEL_TOKEN_ADMIN string
}