package device_core

type DeviceEntityPort interface {
	ShowAllDevices() ([]*DeviceDto, error)
	ShowDeviceByID(id uint) (*DeviceDto, error)
	ShowExpiresDevices(days int) ([]*DeviceDto, error)
	CreateNewDevice(deviceDto *DeviceDto) (*DeviceDto, error)
	ShowDeviceByStation(station string) ([]*DeviceDto, error)
	ShowDeviceByField(field, value string) ([]*DeviceDto, error)
	ChangeDeviceForChecking(device_id, temp_device_id uint) (*DeviceChangeDto, error)
	DeleteDevice(device_id uint) (*DeviceDto, error)
	UploadImage(device_id uint, image_url string) error
}
