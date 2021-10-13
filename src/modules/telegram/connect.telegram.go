package telegram

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	device_core "github.com/hramov/jobhelper/src/core/device"
	user_core "github.com/hramov/jobhelper/src/core/user"
	"github.com/hramov/jobhelper/src/modules/logger"
	device_handler "github.com/hramov/jobhelper/src/modules/telegram/handler/device"
	user_handler "github.com/hramov/jobhelper/src/modules/telegram/handler/user"
	"github.com/hramov/jobhelper/src/modules/telegram/worker"
)

type TGBot struct {
	Instance         *tgbotapi.BotAPI
	Update           tgbotapi.UpdateConfig
	Token            string
	DeviceIDForImage uint
	Admin            string
}

type Users map[int64]*user_core.UserDto

type Message = *tgbotapi.Message
type Bot = *tgbotapi.BotAPI

func (b *TGBot) Create() *TGBot {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		logger.Log("TGBot:Create", err.Error())
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	b.Instance = bot
	b.Update = u

	worker := worker.NotificationWorker{TimePeriod: 10}
	go worker.CheckDevices(b.Instance)

	return b
}

func (b *TGBot) HandleQuery(updateConfig tgbotapi.UpdateConfig) {

	users := make(Users)

	updates, err := b.Instance.GetUpdatesChan(updateConfig)
	if err != nil {
		logger.Log("TGBot:HandleQuery", err.Error())
	}

	logger.Log("TGBot:HandleQuery", "Ready to accept queries!")

	for update := range updates {
		_, ok := users[update.Message.Chat.ID]
		if !ok {
			user, err := user_handler.Check(update.Message.Chat.ID)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Для работы в системе необходима регистрация. Пожалуйста, напишите об этом @%s. Ваш ChatID: %d", b.Admin, update.Message.Chat.ID))
				b.Instance.Send(msg)
				continue
			}
			users[update.Message.Chat.ID] = user
			ok = true
		}

		if update.Message == nil {
			continue
		}

		logger.Log("TGBot:HandleQuery", fmt.Sprintf("New query: %s, ChatID: %d, MessageID: %d", update.Message.Text, update.Message.Chat.ID, update.Message.MessageID))

		if update.Message.Photo != nil {
			var msg tgbotapi.MessageConfig
			err := device_handler.UploadTagImageUrl(b.DeviceIDForImage, (*update.Message.Photo)[3].FileID)
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка: "+err.Error())
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Изображение успешно загружено!")
			}
			b.Instance.Send(msg)
		}

		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			if update.Message.IsCommand() {
				command := update.Message.Command()
				data := update.Message.CommandArguments()

				var deviceReply []*device_core.DeviceDto
				var deviceChangeReply []*device_core.DeviceChangeDto
				var userReply []*user_core.UserDto

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
					deviceReply, err = device_handler.Create(data)
					b.DeviceIDForImage = deviceReply[0].ID
					logger.Log("Create case", "Ready to upload message")
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пожалуйста, загрузите фоторафию бирки для этого оборудования в ответном сообщении")
					b.Instance.Send(msg)
					break
				case "change":
					deviceChangeReply, err = device_handler.Change(data)
					break
				case "myid":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("%d", update.Message.Chat.ID))
					b.Instance.Send(msg)
					break
				case "register":
					_, err = user_handler.Register(data)
					if err != nil {
						logger.Log("TGBot:Handler:Register", err.Error())
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при создании пользователя")
						b.Instance.Send(msg)
						break
					}
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пользователь успешно зарегистрирован")
					b.Instance.Send(msg)
					break
				case "whoami":
					userReply, err = user_handler.WhoAmI(update.Message.Chat.ID)
					break
				case "all":
					deviceReply, err = device_handler.GetAll()
					if len(deviceReply) == 0 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Нет добавленного оборудования")
						b.Instance.Send(msg)
					}
					break
				case "check":
					if data == "" {
						data = "14"
					}
					days, err := strconv.Atoi(data)
					if err != nil {
						break
					}
					deviceReply, err = device_handler.Check(days)
					if len(deviceReply) == 0 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Просроченного оборудования нет!")
						b.Instance.Send(msg)
					}
				case "delete":
					deviceReply, err = device_handler.Delete(data)
					if err == nil {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Успешно удалено следующее оборудование:")
						b.Instance.Send(msg)
					}
					break
				default:
					deviceReply, err = device_handler.GetByField(command, data)
					if len(deviceReply) == 0 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не удалось найти оборудование")
						b.Instance.Send(msg)
					}
					break
				}

				// Displaying error if exists
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
					b.Instance.Send(msg)
					continue
				}

				// Displaying reply
				for _, device := range deviceReply {
					if device.TagImageUrl != "" {
						url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto?chat_id=%d&photo=http://%s:%s/%s", os.Getenv("TOKEN"), update.Message.Chat.ID, os.Getenv("APP_HOST"), os.Getenv("APP_PORT"), device.TagImageUrl)
						_, err := http.Get(url)
						if err != nil {
							logger.Log("Image Sender", err.Error())
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не могу отобразить изображение")
							b.Instance.Send(msg)
							continue
						}
					}

					msg := CreateMessage(*update.Message, fmt.Sprintf("ID: %d\nТип: %s\nНазвание: %s\nОписание: %s\nНомер: %s\nСтанция: %s\nРасположение:%s\nСтатус: %s\nДата проверки: %v\nДата следующей проверки: %v\n\nhttp://%s:%s/%s", device.ID, device.Type, device.Title, device.Description, device.InvNumber, device.Station, device.Location, device.Status, strings.Split(fmt.Sprintf("%s", device.PrevCheck), " ")[0], strings.Split(fmt.Sprintf("%s", device.NextCheck), " ")[0], os.Getenv("APP_HOST"), os.Getenv("APP_PORT"), device.TagImageUrl))
					b.Instance.Send(msg)
				}

				for _, record := range deviceChangeReply {
					msg := CreateMessage(*update.Message, fmt.Sprintf("%d успешно заменен на %d", record.DeviceID, record.TempDeviceID))
					b.Instance.Send(msg)
				}

				for _, user := range userReply {
					msg := CreateMessage(*update.Message, fmt.Sprintf("%s %s\nДолжность: %s\nСтанция: %s\nChat ID: %d", user.LastName, user.Name, user.Position, user.Station, user.ChatID))
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
	format := fmt.Sprintf("Запись оборудования имеет следующие поля:\ntype (Тип);title (Название);(description) Описание;inv_number (Номер блока);station (Станция); location(Место); status(Статус: 'Основной / Подменный');prev_check (Дата проверки (дд.мм.гггг));next_check(Дата следующей проверки (дд.мм.гггг))")
	return fmt.Sprintf("%s\n%s", commands, format)
}
