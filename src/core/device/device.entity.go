package device_core

type DeviceEntity struct {
	Provider       DeviceProviderPort
	ChangeProvider DeviceChangeProviderPort
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

func (d *DeviceEntity) ChangeDeviceForChecking(device_id, temp_device_id uint) (*DeviceChangeDto, error) {
	device, err := d.Provider.FindByID(temp_device_id)
	if err != nil {
		return nil, err
	}
	if device.Status == "В работе" {
		_, err := d.Provider.ReplaceDevice(device_id, temp_device_id)
		if err != nil {
			return nil, err
		}
	} else if device.Status == "На проверке" {
		_, err := d.Provider.ReplaceTempDevice(device_id, temp_device_id)
		if err != nil {
			return nil, err
		}
	}
	record, err := d.ChangeProvider.CreateChangeRecord(device_id, temp_device_id)
	if err != nil {
		return nil, err
	}
	return record, nil
}
