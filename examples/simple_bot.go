package main

import (
	"log"

	"github.com/mattiabrandon/gobot"
)

func messageHandler(bot *gobot.GoBot, message *gobot.Message) {
	if message.Text == "/start" {
		keyboard := [][]*gobot.InlineKeyboardButton{
			{
				{
					Text:         "I'm a button",
					CallbackData: "/callback",
				},
				{
					Text: "This is Google",
					URL:  "https://google.it",
				},
			},
			{
				{
					Text: "This is DuckDuckGo",
					URL:  "https://duckduckgo.com",
				},
			},
		}

		if _, err := bot.SendMessage(gobot.SendMessageParams{
			ChatID:      message.From.ID,
			Text:        "<b>Hello World</b>",
			ReplyMarkup: &gobot.InlineKeyboardMarkup{InlineKeyboard: keyboard},
		}); err != nil {
			log.Println(err)
		}
	} else {
		if _, err := bot.SendMessage(gobot.SendMessageParams{
			ChatID: message.From.ID,
			Text:   "Unrecognised <i>command</i>, try with /start",
		}); err != nil {
			log.Println(err)
		}
	}
}

func callbackQueryHandler(bot *gobot.GoBot, callbackQuery *gobot.CallbackQuery) {
	if callbackQuery.Data == "/back" {
		keyboard := [][]*gobot.InlineKeyboardButton{
			{
				{
					Text:         "I'm a button",
					CallbackData: "/callback",
				},
				{
					Text: "That's Google",
					URL:  "https://google.it",
				},
			},
			{
				{
					Text: "That's DuckDuckGo",
					URL:  "https://duckduckgo.com",
				},
			},
		}

		if _, err := bot.EditMessageText(gobot.EditMessageTextParams{
			ChatID:      callbackQuery.From.ID,
			MessageID:   callbackQuery.Message.MessageID,
			Text:        "<b>Hello World</b>",
			ParseMode:   "HTML",
			ReplyMarkup: &gobot.InlineKeyboardMarkup{InlineKeyboard: keyboard},
		}); err != nil {
			log.Println(err)
		}
	} else if callbackQuery.Data == "/callback" {
		keyboard := [][]*gobot.InlineKeyboardButton{{{
			Text:         "Go back",
			CallbackData: "/back",
		}}}

		if _, err := bot.EditMessageText(gobot.EditMessageTextParams{
			ChatID:      callbackQuery.From.ID,
			MessageID:   callbackQuery.Message.MessageID,
			Text:        "This is a cool callback",
			ParseMode:   "HTML",
			ReplyMarkup: &gobot.InlineKeyboardMarkup{InlineKeyboard: keyboard},
		}); err != nil {
			log.Println(err)
		}
	}
}

func main() {
	bot, err := gobot.Init("TOKEN")

	if err != nil {
		log.Fatalln(err)
	}
	bot.AddHandler(func(bot *gobot.GoBot, update *gobot.Update) {
		if update.Message != nil {
			messageHandler(bot, update.Message)
		} else if update.CallbackQuery != nil {
			callbackQueryHandler(bot, update.CallbackQuery)
		}
	})

	if err = bot.Loop(); err != nil {
		log.Fatalln(err)
	}
}
