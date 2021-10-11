package device_core

type DeviceProviderPort interface {
	Create(device *DeviceDto) (*DeviceDto, error)
	ReplaceDevice(device_id, temp_device_id uint) (*DeviceDto, error)
	ReplaceTempDevice(device_id, temp_device_id uint) (*DeviceDto, error)
	FindAll() ([]*DeviceDto, error)
	FindByID(id uint) (*DeviceDto, error)
	FindByDueDate(days int) ([]*DeviceDto, error)
	FindByStringCondition(field, value string) ([]*DeviceDto, error)
	FindByStation(station string) ([]*DeviceDto, error)
	DeleteDevice(id uint) (*DeviceDto, error)
	SaveDevice(device *DeviceDto) (*DeviceDto, error)
}

type DeviceChangeProviderPort interface {
	CreateChangeRecord(device_id, temp_device_id uint) (*DeviceChangeDto, error)
}
