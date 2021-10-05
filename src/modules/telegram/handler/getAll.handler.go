package handler

import (
	"encoding/json"

	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
	"github.com/hramov/jobhelper/src/modules/logger"
)

func GetAll(message Message, bot Bot) {
	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	devices, err := deviceEntity.ShowAllDevices()
	if err != nil {
		msg := logger.CreateMessage(*message, "Cannot find news :-(")
		bot.Send(msg)
		return
	}

	data, err := json.Marshal(devices)
	msg := logger.CreateMessage(*message, string(data))
	bot.Send(msg)
}
