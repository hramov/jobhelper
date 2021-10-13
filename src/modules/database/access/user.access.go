package access

import (
	"fmt"

	user_core "github.com/hramov/jobhelper/src/core/user"
	"github.com/hramov/jobhelper/src/modules/database/mapper"
	"github.com/hramov/jobhelper/src/modules/database/model"
	"gorm.io/gorm"
)

type UserAccess struct {
	User  *model.User
	Users []*model.User
	DB    *gorm.DB
}

func (ua *UserAccess) FindAll() ([]*user_core.UserDto, error) {
	var users []*user_core.UserDto
	ua.DB.Find(&ua.Users)
	for i := 0; i < len(ua.Users); i++ {
		um := mapper.UserMapper{Model: *ua.Users[i]}
		users = append(users, um.ModelToDto())
	}
	return users, nil
}

func (ua *UserAccess) FindByID(id uint) (*user_core.UserDto, error) {
	ua.User = nil
	ua.DB.Find(&ua.User, "id=?", id)
	um := mapper.UserMapper{Model: *ua.User}
	return um.ModelToDto(), nil
}

func (ua *UserAccess) FindByPosition(position string) ([]*user_core.UserDto, error) {
	ua.User = nil
	var users []*user_core.UserDto
	ua.DB.Find(&ua.Users, "position=?", position)
	for i := 0; i < len(ua.Users); i++ {
		um := mapper.UserMapper{Model: *ua.Users[i]}
		users = append(users, um.ModelToDto())
	}
	return users, nil
}

func (ua *UserAccess) FindByRole(role string) ([]*user_core.UserDto, error) {
	ua.User = nil
	var users []*user_core.UserDto
	ua.DB.Find(&ua.Users, "role=?", role)
	for i := 0; i < len(ua.Users); i++ {
		um := mapper.UserMapper{Model: *ua.Users[i]}
		users = append(users, um.ModelToDto())
	}
	return users, nil
}

func (ua *UserAccess) IsAdmin(id uint) (bool, error) {
	ua.User = nil
	ua.DB.Find(&ua.User, "id=?", id)
	if ua.User.Role == "Админ" {
		return true, nil
	}
	return false, nil
}

func (ua *UserAccess) Create(user *user_core.UserDto) (*user_core.UserDto, error) {
	um := mapper.UserMapper{Dto: *user}
	userModel := um.DtoToModel()
	ua.DB.Create(&userModel)
	result, err := ua.FindByID(userModel.ID)
	if err != nil {
		return nil, fmt.Errorf("Cannot create order")
	}
	return result, nil
}

func (ua *UserAccess) FindByChatID(chat_id int64) (*user_core.UserDto, error) {
	ua.User = nil
	var user *user_core.UserDto
	ua.DB.Find(&ua.User, "chat_id=?", chat_id)
	um := mapper.UserMapper{Model: *ua.User}
	if user = um.ModelToDto(); user.ID == 0 {
		return nil, fmt.Errorf("Пользователь не найден")
	}
	return user, nil
}
