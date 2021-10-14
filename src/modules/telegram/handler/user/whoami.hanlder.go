package user_handler

import (
	"fmt"

	"github.com/golobby/container/v3"
	user_core "github.com/hramov/jobhelper/src/core/user"
)

func WhoAmI(chat_id int64) (*user_core.UserDto, error) {
	var userEntity user_core.UserEntityPort
	container.NamedResolve(&userEntity, "UserEntity")

	user, err := userEntity.ShowUserByChatID(chat_id)
	if err != nil {
		return nil, fmt.Errorf("Не удалось найти пользователя: %s", err.Error())
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("Кажется, вы не зарегистрированы")
	}

	return user, nil
}
