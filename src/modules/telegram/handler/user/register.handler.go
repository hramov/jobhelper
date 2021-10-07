package user_handler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golobby/container/v3"
	user_core "github.com/hramov/jobhelper/src/core/user"
)

func Register(data string) ([]*user_core.UserDto, error) {
	var userEntity user_core.UserEntityPort
	container.NamedResolve(&userEntity, "UserEntity")

	fields := strings.Split(data, " ")
	if len(fields) < 6 {
		return nil, fmt.Errorf("Не хватает данных")
	}

	chat_id, err := strconv.Atoi(fields[5])
	if err != nil {
		return nil, fmt.Errorf("Не удалось создать запись: %s", err.Error())
	}

	userDto := &user_core.UserDto{
		LastName:  fields[0],
		Name:      fields[1],
		Position:  fields[2],
		Role:      fields[3],
		Station:   fields[4],
		ChatID:    int64(chat_id),
		CreatedAt: time.Now(),
	}

	user, err := userEntity.Register(userDto)
	if err != nil {
		return nil, fmt.Errorf("Не удалось создать запись: %s", err.Error())
	}

	return []*user_core.UserDto{user}, nil
}
