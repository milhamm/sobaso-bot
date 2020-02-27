package main

import (
	"fmt"
	"log"
	// "os"
	// "github.com/line/line-bot-sdk-go/linebot"
	"github.com/joho/godotenv"
	"github.com/milhamm/sobaso-bot/config"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main(){
	conf := config.New()

	fmt.Println(conf.Line.CHANNEL_SECRET)
	fmt.Println(conf.Line.CHANNEL_TOKEN)
}