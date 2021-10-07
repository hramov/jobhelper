package access

import (
	"time"

	device_core "github.com/hramov/jobhelper/src/core/device"
	"github.com/hramov/jobhelper/src/modules/database/mapper"
	"github.com/hramov/jobhelper/src/modules/database/model"
	"gorm.io/gorm"
)

type DeviceChangeAccess struct {
	DeviceChangeRecord  *model.DeviceChange
	DeviceChangeRecords []*model.DeviceChange
	DB                  *gorm.DB
}

func (dca *DeviceChangeAccess) CreateChangeRecord(device_id, temp_device_id uint) (*device_core.DeviceChangeDto, error) {
	record := &device_core.DeviceChangeDto{
		DeviceID:     device_id,
		TempDeviceID: temp_device_id,
		ChangedAt:    time.Now(),
	}
	dcm := mapper.DeviceChangeMapper{Dto: *record}
	dcModel := dcm.DtoToModel()
	dca.DB.Create(dcModel)
	return record, nil
}
