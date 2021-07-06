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

	for _, value := range values {
		row = append(row, value)
	}
	return row
}

func NewInlineKeyboardMarkup(values ...[]*InlineKeyboardButton) *InlineKeyboardMarkup {
	keyboard := &InlineKeyboardMarkup{InlineKeyboard: [][]*InlineKeyboardButton{}}

	for _, value := range values {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, value)
	}
	return keyboard
}

func NewSendMessage(chatId int, text string, replyMarkup interface{}) SendMessageParams {
	return SendMessageParams{
		ChatId:      chatId,
		Text:        text,
		ReplyMarkup: replyMarkup,
	}
}

func (message *Message) NewSendMessage(text string, replyMarkup interface{}) SendMessageParams {
	return SendMessageParams{
		ChatId:      message.Chat.Id,
		Text:        text,
		ReplyMarkup: replyMarkup,
	}
}

func NewEditMessageText(chatId int, messageId int, text string, replyMarkup *InlineKeyboardMarkup) EditMessageTextParams {
	return EditMessageTextParams{
		ChatId:      chatId,
		MessageId:   messageId,
		Text:        text,
		ReplyMarkup: replyMarkup,
	}
}

func (message *Message) NewEditMessageText(text string, replyMarkup *InlineKeyboardMarkup) EditMessageTextParams {
	return EditMessageTextParams{
		ChatId:      message.Chat.Id,
		MessageId:   message.MessageId,
		Text:        text,
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
