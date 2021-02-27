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

func (m *Provider) SelectPassengers(passengers *[]*storage.Passenger) error {
	passengerMutex.Lock()
	defer passengerMutex.Unlock()

	for _, passenger := range passengerStorage {
		*passengers = append(*passengers, passenger)
	}

	return nil
}

func (m *Provider) InsertPassenger(passenger *storage.Passenger) error {
	passengerMutex.Lock()
	passengerStorage[passenger.Session.Id] = passenger
	passengerMutex.Unlock()

	drivers := make([]*storage.Driver, 0)
	err := m.SelectDrivers(&drivers)
	if err != nil {
		return err
	}

	for _, driver := range drivers {
		m.updateRoutine(driver, passenger)
	}

	return nil
}

func (m *Provider) UpdatePassenger(passenger *storage.Passenger) error {
	passengerMutex.Lock()
	passengerStorage[passenger.Session.Id] = passenger
	passengerMutex.Unlock()

	drivers := make([]*storage.Driver, 0)
	err := m.SelectDrivers(&drivers)
	if err != nil {
		return err
	}

	for _, driver := range drivers {
		m.updateRoutine(driver, passenger)
	}

	return nil
}

func (m *Provider) DeletePassenger(passenger *storage.Passenger) error {
	passengerMutex.Lock()
	defer passengerMutex.Unlock()

	err := m.deletePassengerAssociatedMappings(passenger.UserId)
	if err != nil {
		return err
	}

	user := storage.NewUser().WithUserId(passenger.UserId)
	err = user.Delete()
	if err != nil {
		return err
	}

	delete(passengerStorage, passenger.Session.Id)
	return nil
}
