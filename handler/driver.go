package handler

import (
	"backend/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserDriver struct {
	User *storage.User
	Driver *storage.Driver
}

type Pickup struct {
	Driver *storage.Driver
	Passenger *storage.Passenger
}

func StartDriverSession(writer http.ResponseWriter, request *http.Request) {
	userDriver := UserDriver{
		User:   storage.NewUser(),
		Driver: storage.NewDriver(),
	}
	err := GetJsonBody(request, userDriver.Driver)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = userDriver.User.WithSessionId(userDriver.Driver.Session.Id).Create()
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = userDriver.Driver.WithUserId(userDriver.User.UserId).Create()
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteJSON(writer, userDriver.Driver.Session)
}

func ConfirmPickup(writer http.ResponseWriter, request *http.Request) {
	sessionId, err := GetSessionId(request)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	pickup := Pickup{
		Driver:    storage.NewDriver(),
		Passenger: storage.NewPassenger(),
	}
	err = pickup.Driver.WithSessionId(sessionId).Select()
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = GetJsonBody(request, pickup.Passenger)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	mapping := storage.NewMapping(pickup.Driver.UserId, pickup.Passenger.UserId)
	err = mapping.Select()
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = pickup.Passenger.Delete()
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = mapping.Delete()
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func GetPassengerInfo(writer http.ResponseWriter, request *http.Request) {
	sessionId, err := GetSessionId(request)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	driver := storage.NewDriver()
	err = driver.WithSessionId(sessionId).Select()
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var mappings []*storage.Mapping
	err = storage.SelectDriverMapping(driver.UserId, &mappings)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var information []*storage.DriverInfo
	for _, mapping := range mappings {
		user := storage.NewUser()
		err = user.WithUserId(mapping.PassengerId.UUId).Select()
		if err != nil {
			logrus.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		driverInfo := &storage.DriverInfo{
			PassengerId: storage.PassengerId{UUId: mapping.PassengerId.UUId},
			Name:        user.Name,
			PickupPoint: storage.LocationLongLat{
				Long: mapping.PickupPoint.Long,
				Lat:  mapping.PickupPoint.Lat,
			},
			DropoffPoint: storage.LocationLongLat{
				Long: mapping.DropoffPoint.Long,
				Lat:  mapping.DropoffPoint.Lat,
			},
			Requested: mapping.Requested,
		}
		information = append(information, driverInfo)
	}
	WriteJSON(writer, information)
}

func UpdateRouteLocations(writer http.ResponseWriter, request *http.Request) {

}

func UpdateEstimations(writer http.ResponseWriter, request *http.Request) {

}

func ConfirmRideRequest(writer http.ResponseWriter, request *http.Request) {

}
