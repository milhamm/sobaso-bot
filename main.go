package main

import (
	"fmt"
	"log"
	// "os"
	// "github.com/line/line-bot-sdk-go/linebot"
	"github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}

func main(){
	conf := NewConfig()

	fmt.Println(conf.Line.CHANNEL_SECRET)
	fmt.Println(conf.Line.CHANNEL_TOKEN)
}