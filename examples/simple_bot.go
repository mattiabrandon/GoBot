package main

import (
	"log"

	"github.com/mattiabrandon/gobot"
)

func messageHandler(bot *gobot.GoBot, update *gobot.Update) {
	if update.Message == nil {
		return
	} else if update.Message.Text == "/start" {
		keyboard := [][]*gobot.InlineKeyboardButton{
			{
				{
					Text:         "I'm a button",
					CallbackData: "/callback",
				},
				{
					Text: "This is Google",
					Url:  "https://google.it",
				},
			},
			{
				{
					Text: "This is DuckDuckGo",
					Url:  "https://duckduckgo.com",
				},
			},
		}

		if _, err := bot.SendMessage(gobot.SendMessageParams{
			ChatId:      update.Message.From.Id,
			Text:        "<b>Hello World</b>",
			ReplyMarkup: &gobot.InlineKeyboardMarkup{InlineKeyboard: keyboard},
		}); err != nil {
			log.Println(err)
		}
	} else {
		if _, err := bot.SendMessage(gobot.SendMessageParams{
			ChatId: update.Message.From.Id,
			Text:   "Unrecognised <i>command</i>, try with /start",
		}); err != nil {
			log.Println(err)
		}
	}
}

func callbackQueryHandler(bot *gobot.GoBot, update *gobot.Update) {
	if update.CallbackQuery.Data == "/back" {
		keyboard := [][]*gobot.InlineKeyboardButton{
			{
				{
					Text:         "I'm a button",
					CallbackData: "/callback",
				},
				{
					Text: "That's Google",
					Url:  "https://google.it",
				},
			},
			{
				{
					Text: "That's DuckDuckGo",
					Url:  "https://duckduckgo.com",
				},
			},
		}

		if _, err := bot.EditMessageText(gobot.EditMessageTextParams{
			ChatId:      update.CallbackQuery.From.Id,
			MessageId:   update.CallbackQuery.Message.MessageId,
			Text:        "<b>Hello World</b>",
			ParseMode:   "HTML",
			ReplyMarkup: &gobot.InlineKeyboardMarkup{InlineKeyboard: keyboard},
		}); err != nil {
			log.Println(err)
		}
	} else if update.CallbackQuery.Data == "/callback" {
		keyboard := [][]*gobot.InlineKeyboardButton{{{
			Text:         "Go back",
			CallbackData: "/back",
		}}}

		if _, err := bot.EditMessageText(gobot.EditMessageTextParams{
			ChatId:      update.CallbackQuery.From.Id,
			MessageId:   update.CallbackQuery.Message.MessageId,
			Text:        "This is a cool callback",
			ParseMode:   "HTML",
			ReplyMarkup: &gobot.InlineKeyboardMarkup{InlineKeyboard: keyboard},
		}); err != nil {
			log.Println(err)
		}
	}
}

func main() {
	bot := gobot.Init("TOKEN")
	bot.AddHandler(gobot.Message{}, messageHandler)
	bot.AddHandler(gobot.CallbackQuery{}, callbackQueryHandler)

	if err := bot.Loop(); err != nil {
		log.Fatalln(err)
	}
}
