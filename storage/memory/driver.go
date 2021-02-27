package memory

import (
	"backend/storage"
	"fmt"
	"sync"
)

var driverStorage map[storage.SessionUUId]*storage.Driver
var driverMutex sync.Mutex

func initDriverStorage() {
	driverStorage = make(map[storage.SessionUUId]*storage.Driver, 0)
}

func (m *Provider) SelectDriver(driver *storage.Driver) error {
	driverMutex.Lock()
	defer driverMutex.Unlock()

	dbDriver, ok := driverStorage[driver.Session.Id]
	if !ok {
		return fmt.Errorf("driver not found")
	}
	*driver = *dbDriver
	return nil
}

func (m *Provider) InsertDriver(driver *storage.Driver) error {
	driverMutex.Lock()
	defer driverMutex.Unlock()

	driverStorage[driver.Session.Id] = driver
	return nil
}

func (m *Provider) UpdateDriver(driver *storage.Driver) error {
	driverMutex.Lock()
	defer driverMutex.Unlock()

	err := m.SelectDriver(driver)
	if err != nil {
		return err
	}
	driverStorage[driver.Session.Id] = driver
	return nil
}

func (m *Provider) DeleteDriver(driver *storage.Driver) error {
	driverMutex.Lock()
	defer driverMutex.Unlock()

	err := m.deleteDriverAssociatedMappings(driver.UserId)
	if err != nil {
		return err
	}
	delete(driverStorage, driver.Session.Id)
	return nil
}
