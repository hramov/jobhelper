package device_core

import (
	"time"
)

type DeviceChangeDto struct {
	ID       uint `json:"id"`
	DeviceID uint `json:"device_id"`
	// Device       *model.Device `json:"device_id"`
	TempDeviceID uint `json:"temp_device_id"`
	// TempDevice   *model.Device `json:"temp_device_id"`
	ChangedAt time.Time `json:"changed_at"`
}
