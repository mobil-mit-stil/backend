package handler

import (
	"backend/storage"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type DebugDump struct {
	Users      []*storage.User      `json:"users"`
	Drivers    []*storage.Driver    `json:"drivers"`
	Passengers []*storage.Passenger `json:"passengers"`
	Mappings   []*storage.Mapping   `json:"mappings"`
}

func NewDebugDump() *DebugDump {
	return &DebugDump{
		Users:      make([]*storage.User, 0),
		Drivers:    make([]*storage.Driver, 0),
		Passengers: make([]*storage.Passenger, 0),
		Mappings:   make([]*storage.Mapping, 0),
	}
}

func DumpDatabase(writer http.ResponseWriter, request *http.Request) {
	dump := NewDebugDump()
	_ = storage.SelectUsers(&dump.Users)
	_ = storage.SelectDrivers(&dump.Drivers)
	_ = storage.SelectPassengers(&dump.Passengers)
	_ = storage.SelectMappings(&dump.Mappings)
	WriteJSON(writer, dump)
}

func CommitNotLive(writer http.ResponseWriter, request *http.Request) {
	logger.Info("i've been shot")
	os.Exit(-1)
}
