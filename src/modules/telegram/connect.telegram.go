package telegram

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	device_core "github.com/hramov/jobhelper/src/core/device"
	"github.com/hramov/jobhelper/src/modules/logger"
	device_handler "github.com/hramov/jobhelper/src/modules/telegram/handler/device"
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
				command := update.Message.Command()
				data := update.Message.CommandArguments()

				fmt.Println(command, data)

				var reply []*device_core.DeviceDto
				var err error

				switch command {
				case "start":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать в систему отслеживания проверки оборудования. Список доступных команд можно посмотреть по /info")
					b.Instance.Send(msg)
					break
				case "info":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, infoMessage())
					b.Instance.Send(msg)
					break
				case "create":
					reply, err = device_handler.Create(data)
					break
				case "all":
					reply, err = device_handler.GetAll()
					break
				case "check":
					days, err := strconv.Atoi(data)
					if err != nil {
						break
					}
					reply, err = device_handler.Check(days)
					if len(reply) == 0 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Просроченного оборудования нет!")
						b.Instance.Send(msg)
					}
				default:
					reply, err = device_handler.GetByField(command, data)
					break
				}

				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
					b.Instance.Send(msg)
					return
				}

				for _, device := range reply {
					msg := logger.CreateMessage(*update.Message, fmt.Sprintf("Тип: %s\nНазвание: %s\nОписание: %s\nНомер: %s\nСтанция: %s\nРасположение:%s\nСтатус: %s\nДата проверки: %v\nДата следующей проверки: %v", device.Type, device.Title, device.Description, device.InvNumber, device.Station, device.Location, device.Status, strings.Split(fmt.Sprintf("%s", device.PrevCheck), " ")[0], strings.Split(fmt.Sprintf("%s", device.NextCheck), " ")[0]))
					b.Instance.Send(msg)
				}

			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неправильный формат комманды. См. /info")
				b.Instance.Send(msg)
			}
		}
	}
}

func infoMessage() string {
	commands := fmt.Sprintf("Список доступных команд:\n/all - посмотреть все оборудование\n/create <Данные оборудования> - записать новое оборудование в базу\n/check <Количество дней до срока проверки> - проверка просроченного оборудования\n/<Поле оборудования> <Значение> - выборка оборудования по определенным полям\n")
	format := fmt.Sprintf("Запись оборудования имеет следующие поля:\ntype (Тип);title (Название);(description) Описание;inv_number (Номер блока);station (Станция); location(Место);prev_check (Дата проверки (дд.мм.гггг));next_check(Дата следующей проверки (дд.мм.гггг))")
	return fmt.Sprintf("%s\n%s", commands, format)
}
