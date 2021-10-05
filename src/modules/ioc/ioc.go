package ioc

import (
	"log"

	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
	device_db "github.com/hramov/jobhelper/src/modules/database/device"
	"gorm.io/gorm"
)

func Register(connection *gorm.DB) error {

	log.Println("Started IoC container!")

	err := container.NamedSingleton("DeviceEntity", func() device_core.DeviceEntityPort {
		return &device_core.DeviceEntity{
			Provider: &device_db.DeviceAccess{
				DB: connection,
			}}
	})

	if err != nil {
		return err
	}
	return nil
}
