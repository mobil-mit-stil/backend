package memory

import (
	"backend/storage"
	"fmt"
)

var passengerStorage map[storage.SessionUUId]*storage.Passenger

func initPassengerStorage() {
	passengerStorage = make(map[storage.SessionUUId]*storage.Passenger, 0)
}

func (m *Provider) SelectPassenger(passenger *storage.Passenger) error {
	dbPassenger, ok := passengerStorage[passenger.SessionId]
	if !ok {
		return fmt.Errorf("passenger not found")
	}
	*passenger = *dbPassenger
	return nil
}

func (m *Provider) InsertPassenger(passenger *storage.Passenger) error {
	passengerStorage[passenger.SessionId] = passenger
	return nil
}

func (m *Provider) UpdatePassenger(passenger *storage.Passenger) error {
	err := m.SelectPassenger(passenger)
	if err != nil {
		return err
	}
	passengerStorage[passenger.SessionId] = passenger
	return nil
}

func (m *Provider) DeletePassenger(passenger *storage.Passenger) error {
	delete(passengerStorage, passenger.SessionId)
	return nil
}
