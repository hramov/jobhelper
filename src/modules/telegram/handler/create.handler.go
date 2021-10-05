package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
	"github.com/hramov/jobhelper/src/modules/logger"
)

func CreateMessage(message Message, bot Bot) {
	msg := logger.CreateMessage(*message, "Пожалуйста, введите через ; следующие поля:")
	bot.Send(msg)
}

func Create(message Message, bot Bot) {

	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	fields := strings.Split(message.Text, ";")
	prevCheck, err := time.Parse("", fields[7])
	nextCheck, err := time.Parse("", fields[8])

	if err != nil {
		msg := logger.CreateMessage(*message, fmt.Sprintf("Не удалось создать запись: %s", err.Error()))
		bot.Send(msg)
		return
	}

	deviceDto := &device_core.DeviceDto{
		Type:        fields[0],
		Title:       fields[1],
		Description: fields[2],
		InvNumber:   fields[3],
		Station:     fields[4],
		Location:    fields[5],
		Status:      fields[6],
		PrevCheck:   prevCheck,
		NextCheck:   nextCheck,
	}

	device, err := deviceEntity.CreateNewDevice(deviceDto)
	if err != nil {
		msg := logger.CreateMessage(*message, fmt.Sprintf("Не удалось создать запись: %s", err.Error()))
		bot.Send(msg)
		return
	}

	msg := logger.CreateMessage(*message, fmt.Sprintf("Успешно записано оборудование: %s", device.Title))
	bot.Send(msg)
}
