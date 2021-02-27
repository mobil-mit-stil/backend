package memory

import (
	"backend/storage"
	"fmt"
)

// map[driver][passenger]
var mappingStorage map[storage.UserUUId]map[storage.UserUUId]*storage.Mapping

func initMappingStorage() {
	mappingStorage = make(map[storage.UserUUId]map[storage.UserUUId]*storage.Mapping, 0)
}

func (m *Provider) SelectSingleMapping(mapping *storage.Mapping) error {
	if !mapping.DriverId.UUId.IsValid() || !mapping.PassengerId.UUId.IsValid() {
		return fmt.Errorf("driverId or passengerId not correct")
	}
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

func (m *Provider) SelectDriverMapping(id storage.UserUUId, mappings *[]*storage.Mapping) error {
	if !id.IsValid() {
		return fmt.Errorf("driverId or passengerId not correct")
	}
	maps, ok := mappingStorage[id]
	if !ok {
		return nil
	}
	for _, mapping := range maps {
		*mappings = append(*mappings, mapping)
	}
	return nil
}

func (m *Provider) SelectPassengerMapping(id storage.UserUUId, mappings *[]*storage.Mapping) error {
	if !id.IsValid() {
		return fmt.Errorf("driverId or passengerId not correct")
	}
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
	ms, ok := mappingStorage[mapping.DriverId.UUId]
	if !ok {
		mappingStorage[mapping.DriverId.UUId] = make(map[storage.UserUUId]*storage.Mapping, 1)
		ms = mappingStorage[mapping.DriverId.UUId]
	}
	delete(ms, mapping.PassengerId.UUId)
	return nil
}
