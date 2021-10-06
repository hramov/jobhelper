package logger

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func CreateMessage(message tgbotapi.Message, text string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	return msg
}
