package device_db

import (
	"fmt"
	"time"

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

func (da *DeviceAccess) FindByID(id uint) (*device_core.DeviceDto, error) {
	da.Device = nil
	da.DB.Find(&da.Device, "id=?", id)
	dm := DeviceMapper{Model: *da.Device}
	return dm.ModelToDto(), nil
}

func (da *DeviceAccess) FindByStation(station string) ([]*device_core.DeviceDto, error) {
	da.Device = nil
	var devices []*device_core.DeviceDto
	da.DB.Find(&da.Devices, "station=?", station)
	for i := 0; i < len(da.Devices); i++ {
		dm := DeviceMapper{Model: *da.Devices[i]}
		devices = append(devices, dm.ModelToDto())
	}
	return devices, nil
}

func (da *DeviceAccess) FindByDueDate(days int) ([]*device_core.DeviceDto, error) {
	da.Device = nil
	var devices []*device_core.DeviceDto
	da.DB.Find(&da.Devices, "next_check < ?", time.Now().AddDate(0, 0, days))
	for i := 0; i < len(da.Devices); i++ {
		dm := DeviceMapper{Model: *da.Devices[i]}
		devices = append(devices, dm.ModelToDto())
	}
	return devices, nil
}

func (da *DeviceAccess) FindByStringCondition(field, value string) ([]*device_core.DeviceDto, error) {
	da.Device = nil
	var devices []*device_core.DeviceDto
	da.DB.Find(&da.Devices, fmt.Sprintf("%s=?", field), value)
	for i := 0; i < len(da.Devices); i++ {
		dm := DeviceMapper{Model: *da.Devices[i]}
		devices = append(devices, dm.ModelToDto())
	}
	return devices, nil
}

func (da *DeviceAccess) Create(device *device_core.DeviceDto) (*device_core.DeviceDto, error) {
	dm := DeviceMapper{Dto: *device}
	deviceModel := dm.DtoToModel()
	da.DB.Create(&deviceModel)
	result, err := da.FindByID(deviceModel.ID)
	if err != nil {
		return nil, fmt.Errorf("Cannot create order")
	}
	return result, nil
}
