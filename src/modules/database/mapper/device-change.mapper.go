package mapper

import (
	device_core "github.com/hramov/jobhelper/src/core/device"
	"github.com/hramov/jobhelper/src/modules/database/model"
)

type DeviceChangeMapper struct {
	Dto   device_core.DeviceChangeDto
	Model model.DeviceChange
}

func (dcm *DeviceChangeMapper) DtoToModel() *model.DeviceChange {
	deviceChange := dcm.Model
	deviceChange.ID = dcm.Dto.ID
	deviceChange.DeviceID = dcm.Dto.DeviceID
	// deviceChange.Device = dcm.Dto.Device
	deviceChange.TempDeviceID = dcm.Dto.TempDeviceID
	// deviceChange.TempDevice = dcm.Dto.TempDevice
	deviceChange.ChangedAt = dcm.Dto.ChangedAt
	return &deviceChange
}

func (dcm *DeviceChangeMapper) ModelToDto() *device_core.DeviceChangeDto {
	deviceChange := dcm.Dto
	deviceChange.ID = dcm.Model.ID
	deviceChange.DeviceID = dcm.Model.DeviceID
	// deviceChange.Device = dcm.Model.Device
	deviceChange.TempDeviceID = dcm.Model.TempDeviceID
	// deviceChange.TempDevice = dcm.Model.TempDevice
	deviceChange.ChangedAt = dcm.Model.ChangedAt
	return &deviceChange
}
