package user_core

type UserEntityPort interface {
	Register(user *UserDto) (*UserDto, error)
	ShowAllUsers() ([]*UserDto, error)
	ShowWhomToSend() ([]*UserDto, error)
}
