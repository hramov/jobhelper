package model

import "time"

type DeviceChange struct {
	ID       uint `json:"id"`
	DeviceID uint `json:"device_id"`
	// Device       *Device   `json:"device_id"`
	TempDeviceID uint `json:"temp_device_id"`
	// TempDevice   *Device   `json:"temp_device_id"`
	ChangedAt time.Time `json:"changed_at"`
}
