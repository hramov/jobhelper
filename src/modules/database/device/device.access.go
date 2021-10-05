package device_db

import (
	device_core "github.com/hramov/jobhelper/src/core/device"
	"gorm.io/gorm"
)

type DeviceAccess struct {
	Device  *Device
	Devices []*Device
	DB      *gorm.DB
}

func (da *DeviceAccess) FindAll() ([]*device_core.DeviceDto, error) {
	da.Device = nil
	var devices []*device_core.DeviceDto
	da.DB.Find(&da.Devices)
	for i := 0; i < len(da.Devices); i++ {
		dm := DeviceMapper{Model: *da.Devices[i]}
		devices = append(devices, dm.ModelToDto())
	}
	return devices, nil
}
