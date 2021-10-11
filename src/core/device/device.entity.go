package device_core

import "fmt"

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

func (d *DeviceEntity) DeleteDevice(id uint) (*DeviceDto, error) {
	return d.Provider.DeleteDevice(id)
}

func (d *DeviceEntity) UploadImage(device_id uint, image_url string) error {

	fmt.Println(device_id, image_url)
	device, err := d.Provider.FindByID(device_id)
	if err != nil {
		return err
	}
	device.TagImageUrl = image_url
	device, err = d.Provider.SaveDevice(device)
	if err != nil {
		return err
	}
	return nil
}
