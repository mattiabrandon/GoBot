package main

import (
	"log"

	"github.com/mattiabrandon/gobot"
)

func main() {
	bot, err := gobot.Init("TOKEN")

	if err != nil {
		log.Fatalln(err)
	}
	bot.AddHandler(func(bot *gobot.GoBot, update *gobot.Update) {
		if update.Message == nil {
			return
		}

		if _, err := bot.SendMessage(gobot.SendMessageParams{
			ChatID:    update.Message.From.ID,
			Text:      update.Message.Text,
			ParseMode: "HTML",
		}); err != nil {
			log.Println(err)
		}
	})

	if err = bot.Loop(); err != nil {
		log.Fatalln(err)
	}
}
