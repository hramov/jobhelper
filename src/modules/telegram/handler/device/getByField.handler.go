package device_handler

import (
	"fmt"

	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
)

func GetByField(field string, value string) ([]*device_core.DeviceDto, error) {
	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	devices, err := deviceEntity.ShowDeviceByField(field, value)
	if err != nil {
		return nil, fmt.Errorf("Не могу найти")
	}
	return devices, nil
}
