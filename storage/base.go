package storage

import (
	"github.com/google/uuid"
)

type SessionUUId string
type UserUUId string

func NewSessionId() SessionUUId {
	return SessionUUId(uuid.New().String())
}

func (s SessionUUId) IsValid() bool {
	_, err := uuid.Parse(string(s))
	return err == nil
}

func NewUserId() UserUUId {
	return UserUUId(uuid.New().String())
}

func (u UserUUId) IsValid() bool {
	_, err := uuid.Parse(string(u))
	return err == nil
}

type Session struct {
	Id SessionUUId `json:"sessionId"`
}

func NewSession() Session {
	return Session{Id: NewSessionId()}
}

type LocationLongLat struct {
	Long float64 `json:"longitude"`
	Lat  float64 `json:"latitude"`
}

type RidePreferences struct {
	Smoker   bool `json:"smoker"`
	Children bool `json:"children"`
}

type DriverId struct {
	UUId UserUUId `json:"driverId"`
}

type PassengerId struct {
	UUId UserUUId `json:"passengerId"`
}
