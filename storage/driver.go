package storage

type Driver struct {
	Session
	UserId      UserUUId          `json:"-"`
	Locations   []LocationLongLat `json:"locations"`
	Seats       int8              `json:"seats"`
	Preferences RidePreferences   `json:"preferences"`
}

func NewDriver() *Driver {
	return &Driver{
		Session: NewSession(),
		UserId: "",
		Locations: make([]LocationLongLat, 0),
		Seats:     0,
		Preferences: RidePreferences{
			Smoker:   false,
			Children: false,
		},
	}
}

func (d *Driver) WithUserId(userId UserUUId) *Driver {
	d.UserId = userId
	return d
}

func (d *Driver) WithLocations(locations *[]LocationLongLat) *Driver {
	d.Locations = *locations
	return d
}

func (d *Driver) WithSeats(seats int8) *Driver {
	d.Seats = seats
	return d
}

func (d *Driver) WithPreferences(preferences *RidePreferences) *Driver {
	d.Preferences = *preferences
	return d
}

func (d *Driver) Select() error {
	return provider.SelectDriver(d)
}

func (d *Driver) Create() error {
	return provider.InsertDriver(d)
}

func (d *Driver) Update() error {
	return provider.UpdateDriver(d)
}

func (d *Driver) Delete() error {
	return provider.DeleteDriver(d)
}
