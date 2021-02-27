package memory

import (
	"backend/storage"
	"fmt"
)

var driverStorage map[storage.SessionUUId]*storage.Driver

func initDriverStorage() {
	driverStorage = make(map[storage.SessionUUId]*storage.Driver, 0)
}

func (m *Provider) SelectDriver(driver *storage.Driver) error {
	dbDriver, ok := driverStorage[driver.SessionId]
	if !ok {
		return fmt.Errorf("driver not found")
	}
	*driver = *dbDriver
	return nil
}

func (m *Provider) InsertDriver(driver *storage.Driver) error {
	driverStorage[driver.SessionId] = driver
	return nil
}

func (m *Provider) UpdateDriver(driver *storage.Driver) error {
	err := m.SelectDriver(driver)
	if err != nil {
		return err
	}
	driverStorage[driver.SessionId] = driver
	return nil
}

func (m *Provider) DeleteDriver(driver *storage.Driver) error {
	delete(driverStorage, driver.SessionId)
	return nil
}
