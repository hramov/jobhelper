package device_handler

import (
	"strconv"
	"strings"

	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
)

func Delete(data string) ([]*device_core.DeviceDto, error) {
	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	idArr := strings.Split(data, " ")
	var devices []*device_core.DeviceDto
	for _, stringId := range idArr {
		id, err := strconv.Atoi(stringId)
		if err != nil {
			return nil, err
		}
		device, err := deviceEntity.DeleteDevice(uint(id))
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}
