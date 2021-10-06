package device_core

type DeviceProviderPort interface {
	Create(device *DeviceDto) (*DeviceDto, error)
	// Update(deviceID uint16, devicekDto *DeviceDto) (*DeviceDto, error)
	// Delete(deviceID uint16) (*DeviceDto, error)
	FindAll() ([]*DeviceDto, error)
	FindByID(id uint) (*DeviceDto, error)
	FindByDueDate(days int) ([]*DeviceDto, error)
	FindByStringCondition(field, value string) ([]*DeviceDto, error)
	// FindByInvNumber(invNumber string) (*DeviceDto, error)
	FindByStation(station string) ([]*DeviceDto, error)
	// FindExpires() (*DeviceDto, error)
}
