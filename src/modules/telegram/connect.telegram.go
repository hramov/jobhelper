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
	"github.com/hramov/jobhelper/src/modules/telegram/middleware"
)

type TGBot struct {
	Instance *tgbotapi.BotAPI
	Update   tgbotapi.UpdateConfig
	Token    string
	Admin    string
	Clients  map[int64]*Client
}

type Client struct {
	User             *user_core.UserDto
	DeviceIDForImage uint
	ErrChan          chan error
}

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
	b.Clients = make(map[int64]*Client)

	// worker := worker.NotificationWorker{TimePeriod: 10}
	// go worker.CheckDevices(b.Instance)

	return b
}

func (b *TGBot) HandleQuery(updateConfig tgbotapi.UpdateConfig) {

	updates, err := b.Instance.GetUpdatesChan(updateConfig)
	if err != nil {
		logger.Log("TGBot:HandleQuery", err.Error())
	}

	logger.Log("TGBot:HandleQuery", "Ready to accept queries!")

	for update := range updates {

		_, ok := b.Clients[update.Message.Chat.ID]
		if !ok {
			user, err := user_handler.Check(update.Message.Chat.ID)
			if err != nil {
				b.sendMessage(update.Message.Chat.ID, fmt.Sprintf("Для работы в системе необходима регистрация. Пожалуйста, напишите об этом @%s. Ваш ChatID: %d", b.Admin, update.Message.Chat.ID))
				continue
			}
			client := &Client{User: user, ErrChan: make(chan error)}
			b.Clients[update.Message.Chat.ID] = client
			ok = true
		}

		if update.Message == nil {
			continue
		}

		logger.Log("TGBot:HandleQuery", fmt.Sprintf("New query: %s, ChatID: %d, MessageID: %d", update.Message.Text, update.Message.Chat.ID, update.Message.MessageID))

		if update.Message.Photo != nil {
			err := device_handler.UploadTagImageUrl(b.Clients[update.Message.Chat.ID].DeviceIDForImage, (*update.Message.Photo)[3].FileID)
			if err != nil {
				b.sendMessage(update.Message.Chat.ID, "Ошибка: "+err.Error())
			} else {
				b.sendMessage(update.Message.Chat.ID, "Изображение успешно загружено!")
			}
		}

		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			if update.Message.IsCommand() {

				go b.commands(update.Message)
				go b.errors(update.Message)

			} else {
				b.sendMessage(update.Message.Chat.ID, "Неправильный формат комманды. См. /info")
			}
		}
	}

}

func (b *TGBot) errors(message *tgbotapi.Message) {
	err := <-b.Clients[message.Chat.ID].ErrChan
	logger.Log("Errors channel", err.Error())
	b.sendMessage(message.Chat.ID, err.Error())
}

func (b *TGBot) sendMessage(chat_id int64, message string) {
	msg := tgbotapi.NewMessage(chat_id, message)
	b.Instance.Send(msg)
}

func (b *TGBot) commands(message *tgbotapi.Message) {

	command := message.Command()
	data := message.CommandArguments()

	chat_id := message.Chat.ID
	client := b.Clients[chat_id]

	perm := middleware.AuthMiddleware(client.User.Role, command)
	if !perm {
		client.ErrChan <- fmt.Errorf("Вам запрещен доступ к этой команде!")
		return
	}

	switch command {
	case "start":
		b.sendMessage(chat_id, "Добро пожаловать в систему отслеживания проверки оборудования. Список доступных команд можно посмотреть по /info")
		break
	case "info":
		commands := fmt.Sprintf("Список доступных команд:\n/all - посмотреть все оборудование\n/create <Данные оборудования> - записать новое оборудование в базу\n/check <Количество дней до срока проверки> - проверка просроченного оборудования\n/<Поле оборудования> <Значение> - выборка оборудования по определенным полям\n")
		format := fmt.Sprintf("Запись оборудования имеет следующие поля:\ntype (Тип);title (Название);(description) Описание;inv_number (Номер блока);station (Станция); location(Место); status(Статус: 'Основной / Подменный');prev_check (Дата проверки (дд.мм.гггг));next_check(Дата следующей проверки (дд.мм.гггг))")
		b.sendMessage(chat_id, fmt.Sprintf("%s\n%s", commands, format))
		break
	case "create":
		result, err := device_handler.Create(data)
		if err != nil {
			client.ErrChan <- err
			return
		}
		b.showDeviceReply(chat_id, result, client.ErrChan)
		client.DeviceIDForImage = result[0].ID
		logger.Log("Create case", "Ready to upload message")
		b.sendMessage(chat_id, "Пожалуйста, загрузите фоторафию бирки для этого оборудования в ответном сообщении")
		break
	case "photo":
		if data != "" {
			result, err := strconv.Atoi(data)
			if err != nil {
				logger.Log("Photo handler", err.Error())
				client.ErrChan <- err
				return
			}
			client.DeviceIDForImage = uint(result)
			b.sendMessage(chat_id, "Пожалуйста, загрузите фоторафию бирки для этого оборудования в ответном сообщении")
		} else {
			b.sendMessage(chat_id, "Пожалуйста, введите команду с идентификатором оборудования")
		}
	case "change":
		result, err := device_handler.Change(data)
		if err != nil {
			client.ErrChan <- err
			return
		}
		b.sendMessage(chat_id, fmt.Sprintf("%d успешно заменен на %d", result.DeviceID, result.TempDeviceID))
		break
	case "myid":
		b.sendMessage(chat_id, fmt.Sprintf("%d", chat_id))
		break
	case "register":
		_, err := user_handler.Register(data)
		if err != nil {
			client.ErrChan <- err
			break
		}
		b.sendMessage(chat_id, "Пользователь успешно зарегистрирован")
		break
	case "whoami":
		result, err := user_handler.WhoAmI(chat_id)
		if err != nil {
			client.ErrChan <- err
			return
		}
		b.sendMessage(chat_id, fmt.Sprintf("%s %s\nДолжность: %s\nСтанция: %s\nChat ID: %d", result.LastName, result.Name, result.Position, result.Station, result.ChatID))
		break
	case "all":
		result, err := device_handler.GetAll()
		if err != nil {
			client.ErrChan <- err
			return
		}
		if len(result) == 0 {
			b.sendMessage(chat_id, "Нет добавленного оборудования")
			return
		}
		b.showDeviceReply(chat_id, result, client.ErrChan)
		break
	case "check":
		if data == "" {
			data = os.Getenv("DEFAULT_CHECK_DAYS")
		}
		days, err := strconv.Atoi(data)
		if err != nil {
			client.ErrChan <- err
			return
		}
		result, err := device_handler.Check(days)
		if len(result) == 0 {
			b.sendMessage(chat_id, "Просроченного оборудования нет!")
		}
	case "delete":
		result, err := device_handler.Delete(data)
		if err == nil {
			client.ErrChan <- err
			return
		}
		b.sendMessage(chat_id, "Успешно удалено следующее оборудование:")
		b.showDeviceReply(chat_id, result, client.ErrChan)
		break
	default:
		result, err := device_handler.GetByField(command, data)
		if err != nil {
			client.ErrChan <- err
			return
		}
		if len(result) == 0 {
			b.sendMessage(chat_id, "Не удалось найти оборудование")
			return
		}
		b.showDeviceReply(chat_id, result, client.ErrChan)
		break
	}
}

func (b *TGBot) showDeviceReply(chat_id int64, devices []*device_core.DeviceDto, errChan chan error) {
	for _, device := range devices {
		if device.TagImageUrl != "" {
			url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto?chat_id=%d&photo=http://%s:%s/%s", os.Getenv("TOKEN"), chat_id, os.Getenv("APP_HOST"), os.Getenv("APP_PORT"), device.TagImageUrl)
			_, err := http.Get(url)
			if err != nil {
				errChan <- err
				logger.Log("Image Sender", err.Error())
				b.sendMessage(chat_id, "Не могу отобразить изображение")
				return
			}
		}
		b.sendMessage(chat_id, fmt.Sprintf("ID: %d\nТип: %s\nНазвание: %s\nОписание: %s\nНомер: %s\nСтанция: %s\nРасположение:%s\nСтатус: %s\nДата проверки: %v\nДата следующей проверки: %v\n\nhttp://%s:%s/%s", device.ID, device.Type, device.Title, device.Description, device.InvNumber, device.Station, device.Location, device.Status, strings.Split(fmt.Sprintf("%s", device.PrevCheck), " ")[0], strings.Split(fmt.Sprintf("%s", device.NextCheck), " ")[0], os.Getenv("APP_HOST"), os.Getenv("APP_PORT"), device.TagImageUrl))
	}
}
