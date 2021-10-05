package device_core

type DeviceEntity struct {
	Provider DeviceProviderPort
}

func (d *DeviceEntity) CreateDevice(device *DeviceDto) (*DeviceDto, error) {
	return d.Provider.Create(device)
}

func (d *DeviceEntity) ShowAllDevices() ([]*DeviceDto, error) {
	return d.Provider.FindAll()
}

func (d *DeviceEntity) ShowDeviceByID(id uint) (*DeviceDto, error) {
	return d.Provider.FindByID(id)
}

func (d *DeviceEntity) ShowExpiresDevices(days int) (*DeviceDto, error) {
	return d.Provider.FindByDueDate(days)
}
