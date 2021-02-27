package memory

import (
	"backend/storage"
	"fmt"
)

var userStorage map[storage.UserUUId]*storage.User

func initUserStorage() {
	userStorage = make(map[storage.UserUUId]*storage.User, 0)
}

func (m *Provider) SelectUser(user *storage.User) error {
	dbUser, ok := userStorage[user.UserId]
	if !ok {
		return fmt.Errorf("user not found")
	}
	*user = *dbUser
	return nil
}

func (m *Provider) InsertUser(user *storage.User) error {
	userStorage[user.UserId] = user
	return nil
}

func (m *Provider) UpdateUser(user *storage.User) error {
	err := m.SelectUser(user)
	if err != nil {
		return err
	}
	userStorage[user.UserId] = user
	return nil
}

func (m *Provider) DeleteUser(user *storage.User) error {
	delete(userStorage, user.UserId)
	return nil
}
