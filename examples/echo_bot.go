package main

import (
	"github.com/mattiabrandon/gobot"
)

func main() {
	bot := gobot.Init("TOKEN")
	bot.AddHandler(gobot.Message{}, func(bot *gobot.GoBot, update *gobot.Update) {
		_, _ = bot.SendMessage(update.Message.NewSendMessage(update.Message.Text, nil))
	})
	_ = bot.Loop(false)
}
