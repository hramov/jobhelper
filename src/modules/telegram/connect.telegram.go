package telegram

import (
	"log"
	"os"
	"reflect"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/hramov/jobhelper/src/modules/telegram/handler"
)

type TGBot struct {
	Instance *tgbotapi.BotAPI
	Update   tgbotapi.UpdateConfig
	Token    string
}

type MessageHistory map[int64]*tgbotapi.Message

type Message = *tgbotapi.Message
type Bot = *tgbotapi.BotAPI

func (b *TGBot) Create() *TGBot {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Println("Error " + err.Error())
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	b.Instance = bot
	b.Update = u
	return b
}

func (b *TGBot) HandleQuery(updateConfig tgbotapi.UpdateConfig) {

	messageHistory := make(MessageHistory)

	updates, err := b.Instance.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Ready to accept queries!")

	for update := range updates {

		log.Printf("New query: %s, MessageID: %d", update.Message.Text, update.Message.MessageID)

		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			messageHistory[update.Message.Chat.ID] = update.Message

			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case "create":
					handler.CreateMessage(update.Message, b.Instance)
				}
			}

			if strings.Contains(update.Message.Text, ";") {
				handler.Create(update.Message, b.Instance)
			}

		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use words for search or command with /.")
			b.Instance.Send(msg)
		}
	}
}
