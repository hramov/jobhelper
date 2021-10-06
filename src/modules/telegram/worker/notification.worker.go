package worker

import (
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
	user_core "github.com/hramov/jobhelper/src/core/user"
)

type NotificationWorker struct {
	TimePeriod int
}

func (nw *NotificationWorker) CheckDevices(bot *tgbotapi.BotAPI) {

	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	var userEntity user_core.UserEntityPort
	container.NamedResolve(&userEntity, "UserEntity")

	now := time.Now()

	for {
		if time.Since(now) > 24*time.Hour {

			log.Println("Notification Worker")

			reply, err := deviceEntity.ShowExpiresDevices(nw.TimePeriod)

			heads, err := userEntity.ShowWhomToSend()
			if err != nil {
				log.Println(err.Error())
			}

			for _, user := range heads {
				for _, device := range reply {
					msg := tgbotapi.NewMessage(user.ChatID, fmt.Sprintf("Тип: %s\nНазвание: %s\nОписание: %s\nНомер: %s\nСтанция: %s\nРасположение:%s\nСтатус: %s\nДата проверки: %v\nДата следующей проверки: %v", device.Type, device.Title, device.Description, device.InvNumber, device.Station, device.Location, device.Status, strings.Split(fmt.Sprintf("%s", device.PrevCheck), " ")[0], strings.Split(fmt.Sprintf("%s", device.NextCheck), " ")[0]))
					bot.Send(msg)
				}
			}

			now = time.Now()
		}
	}
}
