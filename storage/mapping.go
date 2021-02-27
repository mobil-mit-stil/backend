package storage

// /passenger/information response
type PassengerInfo struct {
	DriverId
	Name            string `json:"name"`
	PickupTime      int32  `json:"pickupTime"`
	DestinationTime int32  `json:"destinationTime"`
}

// /driver/estimations request
type Estimation struct {
	PassengerId
	PickupTime      int32 `json:"pickupTime"`
	DestinationTime int32 `json:"destinationTime"`
}

// /driver/confirmations request
type Confirmation struct {
	PassengerId
	Accepted bool `json:"accepted"`
}

// /driver/information response
type DriverInfo struct {
	PassengerId
	Name         string          `json:"name"`
	PickupPoint  LocationLongLat `json:"pickupPoint"`
	DropoffPoint LocationLongLat `json:"dropoffPoint"`
	Requested    bool            `json:"requested"`
}

type Status string

const (
	Accepted Status = "accepted"
	Denied   Status = "denied"
	Pending  Status = "pending"
)

type Mapping struct {
	DriverId
	PassengerId
	PickupPoint     LocationLongLat `json:"pickupPoint"`
	DropoffPoint    LocationLongLat `json:"dropoffPoint"`
	PickupTime      int32           `json:"pickupTime"`
	DestinationTime int32           `json:"destinationTime"`
	Requested       bool            `json:"requested"`
	Status          Status          `json:"accepted"`
}

func NewMapping(driverId UserUUId, passengerId UserUUId) *Mapping {
	return &Mapping{
		DriverId:        DriverId{UUId: driverId},
		PassengerId:     PassengerId{UUId: passengerId},
		PickupPoint:     LocationLongLat{},
		DropoffPoint:    LocationLongLat{},
		PickupTime:      0,
		DestinationTime: 0,
		Requested:       false,
		Status:          "",
	}
}

func (m *Mapping) WithPoints(pickupPoint *LocationLongLat, dropoffPoint *LocationLongLat) *Mapping {
	m.PickupPoint = *pickupPoint
	m.DropoffPoint = *dropoffPoint
	return m
}

func (m *Mapping) WithTimes(pickupTime int32, destinationTime int32) *Mapping {
	m.PickupTime = pickupTime
	m.DestinationTime = destinationTime
	return m
}

func (m *Mapping) WithRequested(requested bool) *Mapping {
	m.Requested = requested
	return m
}

func (m *Mapping) WithStatus(status Status) *Mapping {
	m.Status = status
	return m
}

func (m *Mapping) Select() error {
	return provider.SelectSingleMapping(m)
}

func (m *Mapping) Create() error {
	return provider.InsertMapping(m)
}

func (m *Mapping) Update() error {
	return provider.UpdateMapping(m)
}

func (m *Mapping) Delete() error {
	return provider.DeleteMapping(m)
}

func SelectMappings(mappings *[]*Mapping) error {
	return provider.SelectMappings(mappings)
}

func SelectDriverMapping(id UserUUId, mappings *[]*Mapping) error {
	return provider.SelectDriverMappings(id, mappings)
}

func SelectPassengerMapping(id UserUUId, mappings *[]*Mapping) error {
	return provider.SelectPassengerMappings(id, mappings)
}
