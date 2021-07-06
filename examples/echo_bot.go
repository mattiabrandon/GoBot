package main

import (
	"fmt"
	"github.com/mattiabrandon/gobot"
)

func main() {
	bot := gobot.Init("TOKEN")
	bot.AddHandler(&gobot.Message{}, func(bot *gobot.GoBot, update *gobot.Update) {
		if _, err := bot.SendMessage(update.Message.NewSendMessage(update.Message.Text, nil)); err != nil {
			fmt.Println(err)
		}
	})
	_ = bot.Loop(false)
}
