package device_core

type DeviceEntity struct {
	Provider DeviceProviderPort
}

func (d *DeviceEntity) CreateNewDevice(device *DeviceDto) (*DeviceDto, error) {
	return d.Provider.Create(device)
}

func (d *DeviceEntity) ShowAllDevices() ([]*DeviceDto, error) {
	return d.Provider.FindAll()
}

func (d *DeviceEntity) ShowDeviceByID(id uint) (*DeviceDto, error) {
	return d.Provider.FindByID(id)
}

func (d *DeviceEntity) ShowExpiresDevices(days int) ([]*DeviceDto, error) {
	return d.Provider.FindByDueDate(days)
}

func (d *DeviceEntity) ShowDeviceByStation(station string) ([]*DeviceDto, error) {
	return d.Provider.FindByStation(station)
}

func (d *DeviceEntity) ShowDeviceByField(field, value string) ([]*DeviceDto, error) {
	return d.Provider.FindByStringCondition(field, value)
}
