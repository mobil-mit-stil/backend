package storage

type User struct {
	UserId    UserUUId    `json:"-"`
	SessionId SessionUUId `json:"-"`
	Name      string      `json:"name"`
}

func NewUser() *User {
	return &User{
		UserId:    NewUserId(),
		SessionId: "",
		Name:      "",
	}
}

func (u *User) WithSessionId(sessionId SessionUUId) *User {
	u.SessionId = sessionId
	return u
}

func (u *User) WithName(name string) *User {
	u.Name = name
	return u
}
