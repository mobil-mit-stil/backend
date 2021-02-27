package storage

type DBObject interface {
	Select() error
	Create() error
	Update() error
	Delete() error
}

type Provider interface {
	Init() error

	SelectUser(user *User) error
	InsertUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(user *User) error

	SelectDriver(driver *Driver) error
	InsertDriver(driver *Driver) error
	UpdateDriver(driver *Driver) error
	DeleteDriver(driver *Driver) error

	SelectPassenger(passenger *Passenger) error
	InsertPassenger(passenger *Passenger) error
	UpdatePassenger(passenger *Passenger) error
	DeletePassenger(passenger *Passenger) error

	SelectSingleMapping(mapping *Mapping) error
	SelectDriverMapping(id UserUUId, mappings []*Mapping) error
	SelectPassengerMapping(id UserUUId, mappings []*Mapping) error
	InsertMapping(mapping *Mapping) error
	UpdateMapping(mapping *Mapping) error
	DeleteMapping(mapping *Mapping) error
}

var provider Provider

func Init(specificProvider Provider) error {
	provider = specificProvider
	return provider.Init()
}
