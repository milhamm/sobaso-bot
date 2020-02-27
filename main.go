package main

import (
	"fmt"
	// "os"
	// "github.com/line/line-bot-sdk-go/linebot"
	"github.com/milhamm/sobaso-bot/config"
)

func main(){
	conf := config.New()

	fmt.Println(conf.CHANNEL_SECRET)
	fmt.Println(conf.CHANNEL_TOKEN)
}