package model

type User struct {
	Id       string
	Nickname string
	Name     string
	Email    string
	Avatar   string
}

func (u User) GetRaw() string {
	//TODO implement me
	panic("implement me")
}

func (u User) GetId() string {
	return ""
}

func (u User) GetNickname() string {
	//TODO implement me
	panic("implement me")
}

func (u User) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (u User) GetEmail() string {
	//TODO implement me
	panic("implement me")
}

func (u User) GetAvatar() string {
	//TODO implement me
	panic("implement me")
}

func (u User) GetAccessToken() {
	//TODO implement me
	panic("implement me")
}

func (u User) GetRefreshToken() {
	//TODO implement me
	panic("implement me")
}

func (u User) GetExpiresIn() {
	//TODO implement me
	panic("implement me")
}

func (u User) GetProvider() {
	//TODO implement me
	panic("implement me")
}
