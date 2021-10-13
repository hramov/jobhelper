package user_handler

import (
	"github.com/golobby/container/v3"
	user_core "github.com/hramov/jobhelper/src/core/user"
)

func Check(chat_id int64) (*user_core.UserDto, error) {
	var userEntity user_core.UserEntityPort
	container.NamedResolve(&userEntity, "UserEntity")

	user, err := userEntity.ShowUserByChatID(chat_id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
