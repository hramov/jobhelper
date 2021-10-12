package device_handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
)

func Create(data string) ([]*device_core.DeviceDto, error) {

	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	fields := strings.Split(data, ";")
	if len(fields) < 9 {
		return nil, fmt.Errorf("Не хватает данных")
	}

	prevCheck, err := timePrepare(fields[7])
	nextCheck, err := timePrepare(fields[8])

	if err != nil {
		return nil, fmt.Errorf("Не удалось создать запись: %s", err.Error())
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
		return nil, fmt.Errorf("Не удалось создать запись: %s", err.Error())
	}

	return []*device_core.DeviceDto{device}, nil
}

func timePrepare(rawTime string) (time.Time, error) {
	tArr := strings.Split(rawTime, ".")
	return time.Parse(time.RFC3339, fmt.Sprintf("%s-%s-%sT00:00:00Z", tArr[2], tArr[1], tArr[0]))
}
