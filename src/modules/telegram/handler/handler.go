package handler

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

type Message = *tgbotapi.Message
type Bot = *tgbotapi.BotAPI

func MainSwitch(message Message, bot Bot) {
	switch message.Text {
	case "/create":
		Create(message, bot)
		break
	case "/check":
		Check(message, bot)
		break
	case "/all":
		GetAll(message, bot)
		break
	}
}
