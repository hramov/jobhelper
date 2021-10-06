package user_db

import (
	user_core "github.com/hramov/jobhelper/src/core/user"
)

type UserMapper struct {
	Dto   user_core.UserDto
	Model User
}

func (um *UserMapper) DtoToModel() *User {
	user := um.Model
	user.ID = um.Dto.ID
	user.LastName = um.Dto.LastName
	user.Name = um.Dto.Name
	user.Role = um.Dto.Role
	user.Position = um.Dto.Position
	user.ChatID = um.Dto.ChatID
	user.Station = um.Dto.Station
	user.CreatedAt = um.Dto.CreatedAt
	return &user
}

func (um *UserMapper) ModelToDto() *user_core.UserDto {
	user := um.Dto
	user.ID = um.Model.ID
	user.LastName = um.Model.LastName
	user.Name = um.Model.Name
	user.Role = um.Model.Role
	user.Position = um.Model.Position
	user.ChatID = um.Model.ChatID
	user.Station = um.Model.Station
	user.CreatedAt = um.Model.CreatedAt
	return &user
}
