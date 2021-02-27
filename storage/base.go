package storage

import (
    "github.com/google/uuid"
)

type SessionUUId string
type UserUUId string

func NewSessionId() SessionUUId {
    return SessionUUId(uuid.New().String())
}

func NewUserId() UserUUId {
    return UserUUId(uuid.New().String())
}

type LocationLongLat struct {
    Long float64 `json:"longitude"`
    Lat float64 `json:"latitude"`
}

type RidePreferences struct {
    Smoker bool `json:"smoker"`
    Children bool `json:"children"`
}
