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

func (u *User) WithUserId(userId UserUUId) *User {
	u.UserId = userId
	return u
}

func (u *User) WithName(name string) *User {
	u.Name = name
	return u
}

func (u *User) Select() error {
	return provider.SelectUser(u)
}

func (u *User) Create() error {
	return provider.InsertUser(u)
}

func (u *User) Update() error {
	return provider.UpdateUser(u)
}

func (u *User) Delete() error {
	return provider.DeleteUser(u)
}

func SelectUsers(users *[]*User) error {
	return provider.SelectUsers(users)
}
