package user_core

type UserProviderPort interface {
	Create(device *UserDto) (*UserDto, error)
	FindByPosition(position string) ([]*UserDto, error)
	FindAll() ([]*UserDto, error)
	FindByRole(role string) ([]*UserDto, error)
	FindByChatID(chat_id int64) (*UserDto, error)
}
