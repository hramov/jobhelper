package device_core

type DeviceEntity struct {
	Provider DeviceProviderPort
}

func (d *DeviceEntity) ShowAllDevices() ([]*DeviceDto, error) {
	return d.Provider.FindAll()
}
