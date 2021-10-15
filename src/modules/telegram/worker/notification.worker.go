package worker

import (
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
	user_core "github.com/hramov/jobhelper/src/core/user"
	"github.com/hramov/jobhelper/src/modules/logger"
)

type NotificationWorker struct {
	TimePeriod int
}

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func (nw *NotificationWorker) CheckDevices(bot *tgbotapi.BotAPI) {

	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	var userEntity user_core.UserEntityPort
	container.NamedResolve(&userEntity, "UserEntity")

	logger.Log("Notification Worker", "Started!")

	for {

		t, err := TimeIn(time.Now(), "Local")
		if err != nil {
			logger.Log("Notification worker", "Cannot parse local time")
			return
		}

		if t.Format("15:04") == "08:00" {
			logger.Log("Notification Worker", fmt.Sprintf("Device checking. Time: %v", t.Format("15:04")))
			reply, err := deviceEntity.ShowExpiresDevices(nw.TimePeriod)
			heads, err := userEntity.ShowWhomToSend()
			if err != nil {
				logger.Log("Notification Worker", err.Error())
			}

			for _, user := range heads {
				if reply == nil {
					msg := tgbotapi.NewMessage(user.ChatID, fmt.Sprintf("Просроченного оборудования нет"))
					bot.Send(msg)
					continue
				}
				for _, device := range reply {
					msg := tgbotapi.NewMessage(user.ChatID, fmt.Sprintf("Внимание!\nТип: %s\nНазвание: %s\nОписание: %s\nНомер: %s\nСтанция: %s\nРасположение:%s\nСтатус: %s\nДата проверки: %v\nДата следующей проверки: %v", device.Type, device.Title, device.Description, device.InvNumber, device.Station, device.Location, device.Status, strings.Split(fmt.Sprintf("%s", device.PrevCheck), " ")[0], strings.Split(fmt.Sprintf("%s", device.NextCheck), " ")[0]))
					bot.Send(msg)
				}
			}
		}

		time.Sleep(30 * time.Second)

	}
}
