package storage

type Driver struct {
    SessionId   SessionUUId       `json:"sessionId"`
    UserId      UserUUId          `json:"-"`
    Locations   []LocationLongLat `json:"locations"`
    Seats       int8              `json:"seats"`
    Preferences RidePreferences   `json:"preferences"`
}

func NewDriver() *Driver {
    return &Driver{
        SessionId: NewSessionId(),
        UserId:    NewUserId(),
        Locations: make([]LocationLongLat, 0),
        Seats:     0,
        Preferences: RidePreferences{
            Smoker:   false,
            Children: false,
        },
    }
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

func (d *Driver) Create() error {
    return InsertDriver(d)
}

func (d *Driver) Update() error {
    return UpdateDriver(d)
}

func (d *Driver) Delete() error {
    return DeleteDriver(d)
}
