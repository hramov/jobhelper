package mapper

import (
	device_core "github.com/hramov/jobhelper/src/core/device"
	"github.com/hramov/jobhelper/src/modules/database/model"
)

type DeviceMapper struct {
	Dto   device_core.DeviceDto
	Model model.Device
}

func (dm *DeviceMapper) DtoToModel() *model.Device {
	device := dm.Model
	device.ID = dm.Dto.ID
	device.Type = dm.Dto.Type
	device.Title = dm.Dto.Title
	device.Description = dm.Dto.Description
	device.InvNumber = dm.Dto.InvNumber
	device.Station = dm.Dto.Station
	device.Location = dm.Dto.Location
	device.Status = dm.Dto.Status
	device.PrevCheck = dm.Dto.PrevCheck
	device.NextCheck = dm.Dto.NextCheck
	return &device
}

func (dm *DeviceMapper) ModelToDto() *device_core.DeviceDto {
	device := dm.Dto
	device.ID = dm.Model.ID
	device.Type = dm.Model.Type
	device.Title = dm.Model.Title
	device.Description = dm.Model.Description
	device.InvNumber = dm.Model.InvNumber
	device.Station = dm.Model.Station
	device.Location = dm.Model.Location
	device.Status = dm.Model.Status
	device.PrevCheck = dm.Model.PrevCheck
	device.NextCheck = dm.Model.NextCheck
	return &device
}
