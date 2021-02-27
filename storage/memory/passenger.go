package memory

import (
	"backend/storage"
	"fmt"
	"sync"
)

var passengerStorage map[storage.SessionUUId]*storage.Passenger
var passengerMutex sync.Mutex

func initPassengerStorage() {
	passengerStorage = make(map[storage.SessionUUId]*storage.Passenger, 0)
}

func (m *Provider) SelectPassenger(passenger *storage.Passenger) error {
	passengerMutex.Lock()
	defer passengerMutex.Unlock()

	dbPassenger, ok := passengerStorage[passenger.Session.Id]
	if !ok {
		return fmt.Errorf("passenger not found")
	}
	*passenger = *dbPassenger
	return nil
}

func (m *Provider) InsertPassenger(passenger *storage.Passenger) error {
	passengerMutex.Lock()
	defer passengerMutex.Unlock()

	passengerStorage[passenger.Session.Id] = passenger
	return nil
}

func (m *Provider) UpdatePassenger(passenger *storage.Passenger) error {
	passengerMutex.Lock()
	defer passengerMutex.Unlock()

	err := m.SelectPassenger(passenger)
	if err != nil {
		return err
	}
	passengerStorage[passenger.Session.Id] = passenger
	return nil
}

func (m *Provider) DeletePassenger(passenger *storage.Passenger) error {
	passengerMutex.Lock()
	defer passengerMutex.Unlock()

	err := m.deletePassengerAssociatedMappings(passenger.UserId)
	if err != nil {
		return err
	}
	delete(passengerStorage, passenger.Session.Id)
	return nil
}
