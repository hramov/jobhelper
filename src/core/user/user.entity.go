package user_core

type UserEntity struct {
	Provider UserProviderPort
}

func (ue *UserEntity) Register(user *UserDto) (*UserDto, error) {
	return ue.Provider.Create(user)
}

func (ue *UserEntity) ShowAllUsers() ([]*UserDto, error) {
	return ue.Provider.FindAll()
}

func (ue *UserEntity) ShowWhomToSend() ([]*UserDto, error) {
	heads, err := ue.Provider.FindByPosition("Старший электромеханик")
	if err != nil {
		return nil, err
	}
	admins, err := ue.Provider.FindByRole("Админ")
	if err != nil {
		return nil, err
	}
	heads = append(heads, admins...)
	return heads, nil
}

func (ue *UserEntity) ShowUserByChatID(chat_id int64) (*UserDto, error) {
	return ue.Provider.FindByChatID(chat_id)
}