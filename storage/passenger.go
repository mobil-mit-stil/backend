package storage

type Passenger struct {
	Session
	UserId         UserUUId        `json:"-"`
	Location       LocationLongLat `json:"location"`
	Destination    LocationLongLat `json:"destination"`
	Tolerance      int32           `json:"tolerance"`
	RequestedSeats int8            `json:"requestedSeats"`
	Preferences    RidePreferences `json:"preferences"`
}

func NewPassenger() *Passenger {
	return &Passenger{
		Session: NewSession(),
		UserId:    "",
		Location: LocationLongLat{
			Long: 0,
			Lat:  0,
		},
		Destination: LocationLongLat{
			Long: 0,
			Lat:  0,
		},
		Tolerance:      0,
		RequestedSeats: 0,
		Preferences: RidePreferences{
			Smoker:   false,
			Children: false,
		},
	}
}

func (p *Passenger) WithUserId(userId UserUUId) *Passenger {
	p.UserId = userId
	return p
}

func (p *Passenger) WithLocation(location *LocationLongLat) *Passenger {
	p.Location = *location
	return p
}

func (p *Passenger) WithDestination(destination *LocationLongLat) *Passenger {
	p.Destination = *destination
	return p
}

func (p *Passenger) WithTolerance(tolerance int32) *Passenger {
	p.Tolerance = tolerance
	return p
}

func (p *Passenger) WithSeats(requestedSeats int8) *Passenger {
	p.RequestedSeats = requestedSeats
	return p
}

func (p *Passenger) WithPreferences(preferences *RidePreferences) *Passenger {
	p.Preferences = *preferences
	return p
}

func (p *Passenger) Select() error {
	return provider.SelectPassenger(p)
}

func (p *Passenger) Create() error {
	return provider.InsertPassenger(p)
}

func (p *Passenger) Update() error {
	return provider.UpdatePassenger(p)
}

func (p *Passenger) Delete() error {
	return provider.DeletePassenger(p)
}
