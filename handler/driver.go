package handler

import (
	"backend/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserDriver struct {
	User   *storage.User
	Driver *storage.Driver
}

type Pickup struct {
	Driver    *storage.Driver
	Passenger *storage.Passenger
}

func StartDriverSession(writer http.ResponseWriter, request *http.Request) {
	userDriver := UserDriver{
		User:   storage.NewUser(),
		Driver: storage.NewDriver(),
	}
	err := GetJsonBody(request, userDriver)
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
	sessionId, err := GetSessionId(request)
	if err != nil {
		logrus.Error(err)
		WriteHttpResponse(writer, BadRequest)
		return
	}
	driver := storage.NewDriver()
	err = driver.WithSessionId(sessionId).Select()
	if err != nil {
		logrus.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	var locations []storage.LocationLongLat
	err = GetJsonBody(request, &locations)
	if err != nil {
		logrus.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	err = driver.WithLocations(&locations).Update()
	if err != nil {
		logrus.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	WriteHttpResponse(writer, StatusOk)
}

func UpdateEstimations(writer http.ResponseWriter, request *http.Request) {
	sessionId, err := GetSessionId(request)
	if err != nil {
		logrus.Error(err)
		WriteHttpResponse(writer, BadRequest)
		return
	}
	driver := storage.NewDriver()
	err = driver.WithSessionId(sessionId).Select()
	if err != nil {
		logrus.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	var mappings []*storage.Mapping
	err = storage.SelectDriverMapping(driver.UserId, &mappings)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var estimations []*storage.Estimation
	err = GetJsonBody(request, &estimations)
	for _, mapping := range mappings {
		for _, estimation := range estimations {
			if mapping.PassengerId.UUId == estimation.PassengerId.UUId {
				err = mapping.WithTimes(estimation.PickupTime, estimation.DestinationTime).Update()
				if err != nil {
					logrus.Error(err)
					writer.WriteHeader(http.StatusInternalServerError)
					return
				}
				break
			}
		}
	}
	WriteHttpResponse(writer, StatusOk)
}

func ConfirmRideRequest(writer http.ResponseWriter, request *http.Request) {
	sessionId, err := GetSessionId(request)
	if err != nil {
		logrus.Error(err)
		WriteHttpResponse(writer, BadRequest)
		return
	}
	driver := storage.NewDriver()
	err = driver.WithSessionId(sessionId).Select()
	if err != nil {
		logrus.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	var mappings []*storage.Mapping
	err = storage.SelectDriverMapping(driver.UserId, &mappings)
	if err != nil {
		logrus.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	var confirmations []*storage.Confirmation
	err = GetJsonBody(request, &confirmations)
	for _, mapping := range mappings {
		for _, confirmation := range confirmations {
			if mapping.PassengerId.UUId == confirmation.PassengerId.UUId {
				if confirmation.Accepted {
					err = mapping.WithStatus(storage.Accepted).Update()
				} else {
					err = mapping.WithStatus(storage.Denied).Update()
				}
				if err != nil {
					logrus.Error(err)
					WriteHttpResponse(writer, InternalServerError)
					return
				}
				break
			}
		}
	}
	WriteHttpResponse(writer, StatusOk)
}
