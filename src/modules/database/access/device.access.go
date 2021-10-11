package access

import (
	"fmt"
	"time"

	device_core "github.com/hramov/jobhelper/src/core/device"
	"github.com/hramov/jobhelper/src/modules/database/mapper"
	"github.com/hramov/jobhelper/src/modules/database/model"
	"gorm.io/gorm"
)

type DeviceAccess struct {
	Device  *model.Device
	Devices []*model.Device
	DB      *gorm.DB
}

func (da *DeviceAccess) FindAll() ([]*device_core.DeviceDto, error) {
	da.Device = nil
	var devices []*device_core.DeviceDto
	da.DB.Find(&da.Devices)
	for i := 0; i < len(da.Devices); i++ {
		dm := mapper.DeviceMapper{Model: *da.Devices[i]}
		devices = append(devices, dm.ModelToDto())
	}
	return devices, nil
}

func (da *DeviceAccess) FindByID(id uint) (*device_core.DeviceDto, error) {
	da.Device = nil
	da.DB.Find(&da.Device, "id=?", id)
	dm := mapper.DeviceMapper{Model: *da.Device}
	return dm.ModelToDto(), nil
}

func (da *DeviceAccess) FindByStation(station string) ([]*device_core.DeviceDto, error) {
	da.Device = nil
	var devices []*device_core.DeviceDto
	da.DB.Find(&da.Devices, "station=?", station)
	for i := 0; i < len(da.Devices); i++ {
		dm := mapper.DeviceMapper{Model: *da.Devices[i]}
		devices = append(devices, dm.ModelToDto())
	}
	return devices, nil
}

func (da *DeviceAccess) FindByDueDate(days int) ([]*device_core.DeviceDto, error) {
	da.Device = nil
	var devices []*device_core.DeviceDto
	da.DB.Find(&da.Devices, "next_check < ?", time.Now().AddDate(0, 0, days))
	for i := 0; i < len(da.Devices); i++ {
		dm := mapper.DeviceMapper{Model: *da.Devices[i]}
		devices = append(devices, dm.ModelToDto())
	}
	return devices, nil
}

func (da *DeviceAccess) FindByStringCondition(field, value string) ([]*device_core.DeviceDto, error) {
	da.Device = nil
	var devices []*device_core.DeviceDto
	da.DB.Find(&da.Devices, fmt.Sprintf("%s=?", field), value)
	for i := 0; i < len(da.Devices); i++ {
		dm := mapper.DeviceMapper{Model: *da.Devices[i]}
		devices = append(devices, dm.ModelToDto())
	}
	return devices, nil
}

func (da *DeviceAccess) Create(device *device_core.DeviceDto) (*device_core.DeviceDto, error) {
	dm := mapper.DeviceMapper{Dto: *device}
	deviceModel := dm.DtoToModel()
	da.DB.Create(&deviceModel)
	result, err := da.FindByID(deviceModel.ID)
	if err != nil {
		return nil, fmt.Errorf("Cannot create order")
	}
	return result, nil
}

func (da *DeviceAccess) DeleteDevice(id uint) (*device_core.DeviceDto, error) {
	da.Device = nil
	device, err := da.FindByID(id)
	if err != nil {
		return nil, err
	}
	if device.ID != 0 {
		da.DB.Delete(&da.Device, "id=?", id)
		return device, nil
	}
	return nil, fmt.Errorf("Оборудование с ID %d не найдено", id)
}

func (da *DeviceAccess) ReplaceDevice(device_id, temp_device_id uint) (*device_core.DeviceDto, error) {
	device, err := da.FindByID(device_id)
	if err != nil {
		return nil, err
	}
	device.Status = "На проверке"
	dm := mapper.DeviceMapper{Dto: *device}
	deviceModel := dm.DtoToModel()
	da.DB.Save(deviceModel)
	return device, nil
}

func (da *DeviceAccess) ReplaceTempDevice(temp_device_id, device_id uint) (*device_core.DeviceDto, error) {
	device, err := da.FindByID(device_id)
	if err != nil {
		return nil, err
	}
	device.Status = "В работе"
	dm := mapper.DeviceMapper{Dto: *device}
	deviceModel := dm.DtoToModel()
	da.DB.Save(deviceModel)
	_, err = da.DeleteDevice(temp_device_id)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (da *DeviceAccess) SaveDevice(device *device_core.DeviceDto) (*device_core.DeviceDto, error) {
	dm := mapper.DeviceMapper{Dto: *device}
	deviceModel := dm.DtoToModel()
	da.DB.Save(deviceModel)
	return device, nil
}
