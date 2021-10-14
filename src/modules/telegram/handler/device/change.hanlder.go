package device_handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
)

func Change(data string) (*device_core.DeviceChangeDto, error) {
	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	ids := strings.Split(data, " ")
	if len(ids) < 2 {
		return nil, fmt.Errorf("Не хватает данных!")
	}

	device_id, err := strconv.Atoi(ids[0])
	if err != nil {
		return nil, err
	}

	temp_device_id, err := strconv.Atoi(ids[1])
	if err != nil {
		return nil, err
	}

	record, err := deviceEntity.ChangeDeviceForChecking(uint(device_id), uint(temp_device_id))
	if err != nil {
		return nil, err
	}

	return record, nil
}
