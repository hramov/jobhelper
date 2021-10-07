package device_handler

import (
	"fmt"

	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
)

func GetAll() ([]*device_core.DeviceDto, error) {
	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	devices, err := deviceEntity.ShowAllDevices()
	if err != nil {
		return nil, fmt.Errorf("Не могу найти: %s", err.Error())
	}
	return devices, nil
}
