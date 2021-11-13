package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/A1esandr/tgbotapi"
)

//go:embed token.txt
var embedToken string

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		token = embedToken
	}
	if token == "" {
		log.Panic("token is empty!")
	}
	bot, err := tgbotapi.New(token)
	if err != nil {
		log.Fatal(err)
	}
	offset := 0
	for {
		resp, err := bot.GetUpdates(&tgbotapi.GetUpdates{
			Offset:  offset,
			Limit:   10,
			Timeout: 1,
		})
		if err != nil {
			log.Fatal(err)
		}
		if resp != nil {
			for _, upd := range resp.Result {
				fmt.Println("Update ID", upd.UpdateID)
				fmt.Println("Chat ID", upd.ChannelPost.Chat.ID)
				fmt.Println("Chat", upd.ChannelPost.Chat.Title)
			}
			if len(resp.Result) < 10 {
				break
			}
			offset += 10
		} else {
			break
		}
	}
	fmt.Println("finish")
}
