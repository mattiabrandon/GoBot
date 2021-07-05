package main

import (
	"log"

	"github.com/mattiabrandon/gobot"
)

func main() {
	bot := gobot.Init("TOKEN")
	bot.AddHandler(gobot.Message{}, func(bot *gobot.GoBot, update *gobot.Update) {
		if _, err := bot.SendMessage(gobot.SendMessageParams{
			ChatId:    update.Message.From.Id,
			Text:      update.Message.Text,
			ParseMode: "HTML",
		}); err != nil {
			log.Println(err)
		}
	})

	if err := bot.Loop(); err != nil {
		log.Fatalln(err)
	}
}
