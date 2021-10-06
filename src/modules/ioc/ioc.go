package ioc

import (
	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
	user_core "github.com/hramov/jobhelper/src/core/user"
	device_db "github.com/hramov/jobhelper/src/modules/database/device"
	user_db "github.com/hramov/jobhelper/src/modules/database/user"
	"github.com/hramov/jobhelper/src/modules/logger"
	"gorm.io/gorm"
)

func Register(connection *gorm.DB) error {

	logger.Log("IoC", "Started IoC container!")

	err := container.NamedSingleton("DeviceEntity", func() device_core.DeviceEntityPort {
		return &device_core.DeviceEntity{
			Provider: &device_db.DeviceAccess{
				DB: connection,
			}}
	})

	err = container.NamedSingleton("UserEntity", func() user_core.UserEntityPort {
		return &user_core.UserEntity{
			Provider: &user_db.UserAccess{
				DB: connection,
			}}
	})

	if err != nil {
		return err
	}
	return nil
}
