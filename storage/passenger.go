package storage

type Passenger struct {
    SessionId      SessionUUId     `json:"sessionId"`
    UserId         UserUUId        `json:"-"`
    Location       LocationLongLat `json:"location"`
    Destination    LocationLongLat `json:"destination"`
    Tolerance      int32           `json:"tolerance"`
    RequestedSeats int8            `json:"requestedSeats"`
    Preferences    RidePreferences `json:"preferences"`
}

func NewPassenger() *Passenger {
    return &Passenger{
        SessionId: NewSessionId(),
        UserId:    NewUserId(),
        Location:       LocationLongLat{
            Long: 0,
            Lat:  0,
        },
        Destination:    LocationLongLat{
            Long: 0,
            Lat:  0,
        },
        Tolerance:      0,
        RequestedSeats: 0,
        Preferences:    RidePreferences{
            Smoker:   false,
            Children: false,
        },
    }
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

func (p *Passenger) Create() error {
    return InsertPassenger(p)
}

func (p *Passenger) Update() error {
    return UpdatePassenger(p)
}

func (p *Passenger) Delete() error {
    return DeletePassenger(p)
}
