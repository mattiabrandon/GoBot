package gobot

func NewInlineKeyboardButton(text string, value string, isCallbackData bool) *InlineKeyboardButton {
	if isCallbackData {
		return &InlineKeyboardButton{
			Text:         text,
			CallbackData: value,
		}
	}
	return &InlineKeyboardButton{
		Text: text,
		Url:  value,
	}
}

func NewInlineKeyboardRow(values ...*InlineKeyboardButton) []*InlineKeyboardButton {
	var row []*InlineKeyboardButton
	row = append(row, values...)
	return row
}

func NewInlineKeyboardMarkup(values ...[]*InlineKeyboardButton) *InlineKeyboardMarkup {
	keyboard := &InlineKeyboardMarkup{InlineKeyboard: [][]*InlineKeyboardButton{}}
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, values...)
	return keyboard
}

func NewSendMessage(chatId int, text string, replyMarkup interface{}) SendMessageParams {
	return SendMessageParams{
		ChatId:      chatId,
		Text:        text,
		ParseMode:   "HTML",
		ReplyMarkup: replyMarkup,
	}
}

func (message *Message) NewSendMessage(text string, replyMarkup interface{}) SendMessageParams {
	return SendMessageParams{
		ChatId:      message.Chat.Id,
		Text:        text,
		ParseMode:   "HTML",
		ReplyMarkup: replyMarkup,
	}
}

func NewEditMessageText(chatId int, messageId int, text string, replyMarkup *InlineKeyboardMarkup) EditMessageTextParams {
	return EditMessageTextParams{
		ChatId:      chatId,
		MessageId:   messageId,
		Text:        text,
		ParseMode:   "HTML",
		ReplyMarkup: replyMarkup,
	}
}

func (message *Message) NewEditMessageText(text string, replyMarkup *InlineKeyboardMarkup) EditMessageTextParams {
	return EditMessageTextParams{
		ChatId:      message.Chat.Id,
		MessageId:   message.MessageId,
		Text:        text,
		ParseMode:   "HTML",
		ReplyMarkup: replyMarkup,
	}
}

func NewAnswerCallbackQuery(callbackQueryId string, text string, showAlert bool) AnswerCallbackQueryParams {
	return AnswerCallbackQueryParams{
		CallbackQueryId: callbackQueryId,
		Text:            text,
		ShowAlert:       showAlert,
	}
}

func (callbackQuery *CallbackQuery) NewAnswerCallbackQuery(text string, showAlert bool) AnswerCallbackQueryParams {
	return AnswerCallbackQueryParams{
		CallbackQueryId: callbackQuery.Id,
		Text:            text,
		ShowAlert:       showAlert,
	}
}
