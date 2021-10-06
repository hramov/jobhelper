package device_handler

import (
	"fmt"

	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
)

func Check(days int) ([]*device_core.DeviceDto, error) {
	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	devices, err := deviceEntity.ShowExpiresDevices(days)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return devices, nil
}
