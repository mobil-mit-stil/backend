package storage

type DBObject interface {
	Select() error
	Create() error
	Update() error
	Delete() error
}

type Provider interface {
	Init() error

	SelectUsers(users *[]*User) error
	SelectUser(user *User) error
	InsertUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(user *User) error

	SelectDrivers(drivers *[]*Driver) error
	SelectDriver(driver *Driver) error
	InsertDriver(driver *Driver) error
	UpdateDriver(driver *Driver) error
	DeleteDriver(driver *Driver) error

	SelectPassengers(passengers *[]*Passenger) error
	SelectPassenger(passenger *Passenger) error
	InsertPassenger(passenger *Passenger) error
	UpdatePassenger(passenger *Passenger) error
	DeletePassenger(passenger *Passenger) error

	SelectMappings(mappings *[]*Mapping) error
	SelectSingleMapping(mapping *Mapping) error
	SelectDriverMappings(id UserUUId, mappings *[]*Mapping) error
	SelectPassengerMappings(id UserUUId, mappings *[]*Mapping) error
	InsertMapping(mapping *Mapping) error
	UpdateMapping(mapping *Mapping) error
	DeleteMapping(mapping *Mapping) error
}

var provider Provider

func Init(specificProvider Provider) error {
	provider = specificProvider
	return provider.Init()
}
