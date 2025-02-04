package memory

import (
	"backend/algorithm"
	"backend/storage"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"sync"
)

// map[driver][passenger]
var mappingStorage map[storage.UserUUId]map[storage.UserUUId]*storage.Mapping
var mappingMutex sync.Mutex

func initMappingStorage() {
	mappingStorage = make(map[storage.UserUUId]map[storage.UserUUId]*storage.Mapping, 0)
}

func (m *Provider) SelectMappings(mappings *[]*storage.Mapping) error {
	mappingMutex.Lock()
	defer mappingMutex.Unlock()
	for _, driverMappings := range mappingStorage {
		for _, mapping := range driverMappings {
			*mappings = append(*mappings, mapping)
		}
	}
	return nil
}

func (m *Provider) SelectSingleMapping(mapping *storage.Mapping) error {
	if !mapping.DriverId.UUId.IsValid() || !mapping.PassengerId.UUId.IsValid() {
		return fmt.Errorf("driverId or passengerId not correct")
	}
	mappingMutex.Lock()
	defer mappingMutex.Unlock()

	maps, ok := mappingStorage[mapping.DriverId.UUId]
	if !ok {
		return fmt.Errorf("could not find mapping for driver")
	}
	for _, candidate := range maps {
		if candidate.PassengerId.UUId == mapping.PassengerId.UUId {
			*mapping = *candidate
			return nil
		}
	}
	return fmt.Errorf("could not find mapping for passenger")
}

func (m *Provider) SelectDriverMappings(id storage.UserUUId, mappings *[]*storage.Mapping) error {
	if !id.IsValid() {
		return fmt.Errorf("driverId or passengerId not correct")
	}
	mappingMutex.Lock()
	defer mappingMutex.Unlock()

	maps, ok := mappingStorage[id]
	if !ok {
		return nil
	}
	for _, mapping := range maps {
		*mappings = append(*mappings, mapping)
	}
	return nil
}

func (m *Provider) SelectPassengerMappings(id storage.UserUUId, mappings *[]*storage.Mapping) error {
	if !id.IsValid() {
		return fmt.Errorf("driverId or passengerId not correct")
	}
	mappingMutex.Lock()
	defer mappingMutex.Unlock()

	for _, maps := range mappingStorage {
		for _, mapping := range maps {
			if mapping.PassengerId.UUId == id {
				*mappings = append(*mappings, mapping)
			}
		}
	}
	return nil
}

func (m *Provider) InsertMapping(mapping *storage.Mapping) error {
	if !mapping.DriverId.UUId.IsValid() || !mapping.PassengerId.UUId.IsValid() {
		return fmt.Errorf("driverId or passengerId not correct")
	}
	mappingMutex.Lock()
	defer mappingMutex.Unlock()

	ms, ok := mappingStorage[mapping.DriverId.UUId]
	if !ok {
		mappingStorage[mapping.DriverId.UUId] = make(map[storage.UserUUId]*storage.Mapping, 1)
		ms = mappingStorage[mapping.DriverId.UUId]
	}
	_, ok = ms[mapping.PassengerId.UUId]
	if ok {
		return fmt.Errorf("mapping already exists")
	}
	ms[mapping.PassengerId.UUId] = mapping
	return nil
}

func (m *Provider) UpdateMapping(mapping *storage.Mapping) error {
	if !mapping.DriverId.UUId.IsValid() || !mapping.PassengerId.UUId.IsValid() {
		return fmt.Errorf("driverId or passengerId not correct")
	}
	mappingMutex.Lock()
	defer mappingMutex.Unlock()

	ms, ok := mappingStorage[mapping.DriverId.UUId]
	if !ok {
		mappingStorage[mapping.DriverId.UUId] = make(map[storage.UserUUId]*storage.Mapping, 1)
		ms = mappingStorage[mapping.DriverId.UUId]
	}
	ms[mapping.PassengerId.UUId] = mapping
	return nil
}

func (m *Provider) DeleteMapping(mapping *storage.Mapping) error {
	if !mapping.DriverId.UUId.IsValid() || !mapping.PassengerId.UUId.IsValid() {
		return fmt.Errorf("driverId or passengerId not correct")
	}
	mappingMutex.Lock()
	defer mappingMutex.Unlock()

	ms, ok := mappingStorage[mapping.DriverId.UUId]
	if !ok {
		mappingStorage[mapping.DriverId.UUId] = make(map[storage.UserUUId]*storage.Mapping, 1)
		ms = mappingStorage[mapping.DriverId.UUId]
	}
	delete(ms, mapping.PassengerId.UUId)
	return nil
}

func (m *Provider) deletePassengerAssociatedMappings(id storage.UserUUId) error {
	mappings := make([]*storage.Mapping, 0)
	err := m.SelectPassengerMappings(id, &mappings)
	if err != nil {
		return err
	}
	for _, mapping := range mappings {
		err = m.DeleteMapping(mapping)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Provider) deleteDriverAssociatedMappings(id storage.UserUUId) error {
	mappings := make([]*storage.Mapping, 0)
	err := m.SelectDriverMappings(id, &mappings)
	if err != nil {
		return err
	}
	for _, mapping := range mappings {
		err = m.DeleteMapping(mapping)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Provider) updateRoutine(driver *storage.Driver, passenger *storage.Passenger) {
	startPoint, ok := algorithm.NearestTolerablePoint(passenger.Location, driver.Locations, passenger.Tolerance)
	if !ok {
		return
	}
	endPoint, ok := algorithm.NearestTolerablePoint(passenger.Destination, driver.Locations, passenger.Tolerance)
	if !ok {
		return
	}
	mapping := storage.NewMapping(driver.UserId, passenger.UserId)
	err := mapping.Select()

	mapping.WithPoints(&startPoint, &endPoint)
	// err is nil when a mapping already exists and therefore needs to be updated, otherwise create new one
	if err == nil {
		err = mapping.Update()
	} else {
		err = mapping.Create()
	}

	if err != nil {
		logger.Warn(err)
	}
}
