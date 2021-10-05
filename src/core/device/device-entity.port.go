package device_core

type DeviceEntityPort interface {
	ShowAllDevices() ([]*DeviceDto, error)
	ShowDeviceByID(id uint) (*DeviceDto, error)
	ShowExpiresDevices(days int) ([]*DeviceDto, error)
	CreateNewDevice(deviceDto *DeviceDto) (*DeviceDto, error)
	// Create(devicekDto *DeviceDto) (*DeviceDto, error)
	// Update(deviceID uint16, devicekDto *DeviceDto) (*DeviceDto, error)
	// Delete(deviceID uint16) (*DeviceDto, error)
	// FindAll() ([]*DeviceDto, error)
	// FindById(id uint64) (*DeviceDto, error)
	// FindByInvNumber(invNumber string) (*DeviceDto, error)
	// FindByStation(station string) (*DeviceDto, error)
	// FindExpires() (*DeviceDto, error)
}
