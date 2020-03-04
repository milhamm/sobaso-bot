package main

import (
	"fmt"
	// "os"
	"github.com/milhamm/sobaso-bot/config"
	"net/http"
	"log"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
)

func main(){
	conf := config.New()

	fmt.Println(conf.CHANNEL_SECRET_CLIENT)
	fmt.Println(conf.CHANNEL_TOKEN_CLIENT)

	client, err := linebot.New(
		conf.CHANNEL_SECRET_CLIENT,
		conf.CHANNEL_TOKEN_CLIENT,
	)

	if err != nil {
		log.Fatal(err)
	}

	admin, err := linebot.New(
		conf.CHANNEL_SECRET_ADMIN,
		conf.CHANNEL_TOKEN_ADMIN,
	)

	http.Handle("/callback", )

	if err := http.ListenAndServe(":5000",nil); err != nil{
		log.Fatal(err)
	}
}