package main

import (
	"github.com/mattiabrandon/gobot"
)

func messageHandler(bot *gobot.GoBot, update *gobot.Update) {
	message := update.Message

	if message.Text == "/start" {
		keyboard := gobot.NewInlineKeyboardMarkup(
			gobot.NewInlineKeyboardRow(
				gobot.NewInlineKeyboardButton("I'm a button", "/callback", true),
				gobot.NewInlineKeyboardButton("This is Google", "https://google.it", false),
			),
			gobot.NewInlineKeyboardRow(
				gobot.NewInlineKeyboardButton("This is DuckDuckGo", "https://duckduckgo.com", false),
			),
		)
		_, _ = bot.SendMessage(message.NewSendMessage(
			"<b>Hello World</b>",
			keyboard,
		))
	} else {
		_, _ = bot.SendMessage(message.NewSendMessage(
			"Unrecognised <i>command</i>, try with /start",
			nil,
		))
	}
}

func callbackQueryHandler(bot *gobot.GoBot, update *gobot.Update) {
	callbackQuery := update.CallbackQuery

	if callbackQuery.Data == "/back" {
		keyboard := gobot.NewInlineKeyboardMarkup(
			gobot.NewInlineKeyboardRow(
				gobot.NewInlineKeyboardButton("I'm a button", "/callback", true),
				gobot.NewInlineKeyboardButton("This is Google", "https://google.it", false),
			),
			gobot.NewInlineKeyboardRow(
				gobot.NewInlineKeyboardButton("This is DuckDuckGo", "https://duckduckgo.com", false),
			),
		)
		_, _ = bot.EditMessageText(callbackQuery.Message.NewEditMessageText(
			"<b>Hello World</b>",
			keyboard,
		))
	} else if update.CallbackQuery.Data == "/callback" {
		keyboard := gobot.NewInlineKeyboardMarkup(gobot.NewInlineKeyboardRow(gobot.NewInlineKeyboardButton(
			"Go back",
			"/back",
			true,
		)))
		_, _ = bot.EditMessageText(callbackQuery.Message.NewEditMessageText(
			"This is a cool callback",
			keyboard,
		))
	}
	_, _ = bot.AnswerCallbackQuery(callbackQuery.NewAnswerCallbackQuery("", false))
}

func main() {
	bot := gobot.Init("TOKEN")
	bot.AddHandler(gobot.Message{}, messageHandler)
	bot.AddHandler(gobot.CallbackQuery{}, callbackQueryHandler)
	_ = bot.Loop(false)
}
