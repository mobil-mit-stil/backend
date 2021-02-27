package storage

type DBObject interface {
    Create() error
    Update() error
    Delete() error
}

type Provider interface {
    Init() error
    InsertDriver(driver *Driver) error
    UpdateDriver(driver *Driver) error
    DeleteDriver(driver *Driver) error
    InsertPassenger(passenger *Passenger) error
    UpdatePassenger(passenger *Passenger) error
    DeletePassenger(passenger *Passenger) error
}

var provider Provider

func Init(specificProvider Provider) error {
    provider = specificProvider
    return provider.Init()
}

func InsertDriver(driver *Driver) error {
    return provider.InsertDriver(driver)
}

func UpdateDriver(driver *Driver) error {
    return provider.UpdateDriver(driver)
}

func DeleteDriver(driver *Driver) error {
    return provider.DeleteDriver(driver)
}

func InsertPassenger(passenger *Passenger) error {
    return provider.InsertPassenger(passenger)
}

func UpdatePassenger(passenger *Passenger) error {
    return provider.UpdatePassenger(passenger)
}

func DeletePassenger(passenger *Passenger) error {
    return provider.DeletePassenger(passenger)
}
