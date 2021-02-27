package memory

import (
	"backend/storage"
	"fmt"
	"sync"
)

var userStorage map[storage.UserUUId]*storage.User
var userMutex sync.Mutex

func initUserStorage() {
	userStorage = make(map[storage.UserUUId]*storage.User, 0)
}

func (m *Provider) SelectUsers(users *[]*storage.User) error {
	userMutex.Lock()
	defer userMutex.Unlock()
	for _, user := range userStorage {
		*users = append(*users, user)
	}
	return nil
}

func (m *Provider) SelectUser(user *storage.User) error {
	userMutex.Lock()
	defer userMutex.Unlock()
	dbUser, ok := userStorage[user.UserId]
	if !ok {
		return fmt.Errorf("user not found")
	}
	*user = *dbUser
	return nil
}

func (m *Provider) InsertUser(user *storage.User) error {
	userMutex.Lock()
	defer userMutex.Unlock()
	userStorage[user.UserId] = user
	return nil
}

func (m *Provider) UpdateUser(user *storage.User) error {
	userMutex.Lock()
	defer userMutex.Unlock()
	err := m.SelectUser(user)
	if err != nil {
		return err
	}
	userStorage[user.UserId] = user
	return nil
}

func (m *Provider) DeleteUser(user *storage.User) error {
	userMutex.Lock()
	defer userMutex.Unlock()
	delete(userStorage, user.UserId)
	return nil
}
