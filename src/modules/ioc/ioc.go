package ioc

import (
	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
	team_core "github.com/hramov/jobhelper/src/core/team"
	user_core "github.com/hramov/jobhelper/src/core/user"
	"github.com/hramov/jobhelper/src/modules/database/access"
	"github.com/hramov/jobhelper/src/modules/logger"
	"gorm.io/gorm"
)

func Register(connection *gorm.DB) error {

	logger.Log("IoC", "Started IoC container!")

	err := container.NamedSingleton("DeviceEntity", func() device_core.DeviceEntityPort {
		return &device_core.DeviceEntity{
			Provider: &access.DeviceAccess{
				DB: connection,
			},
			ChangeProvider: &access.DeviceChangeAccess{
				DB: connection,
			}}
	})

	err = container.NamedSingleton("UserEntity", func() user_core.UserEntityPort {
		return &user_core.UserEntity{
			Provider: &access.UserAccess{
				DB: connection,
			}}
	})

	err = container.NamedSingleton("TeamEntity", func() team_core.TeamEntityPort {
		return &team_core.TeamEntity{
			Provider: &access.TeamAccess{
				DB: connection,
			}}
	})

	if err != nil {
		return err
	}
	return nil
}
