package handler

import (
	"backend/storage"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

type UserDriver struct {
	*storage.User
	*storage.Driver
}

type Pickup struct {
	*storage.Driver
	*storage.PassengerId
}

func StartDriverSession(writer http.ResponseWriter, request *http.Request) {
	userDriver := &UserDriver{
		User:   storage.NewUser(),
		Driver: storage.NewDriver(),
	}
	err := GetJsonBody(request, userDriver)
	if err != nil {
		logger.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = userDriver.User.WithSessionId(userDriver.Driver.Session.Id).Create()
	if err != nil {
		logger.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = userDriver.Driver.WithUserId(userDriver.User.UserId).Create()
	if err != nil {
		logger.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteJSON(writer, userDriver.Driver.Session)
}

func ConfirmPickup(writer http.ResponseWriter, request *http.Request) {
	sessionId, err := GetSessionId(request)
	if err != nil {
		logger.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	pickup := &Pickup{
		Driver:      storage.NewDriver(),
		PassengerId: &storage.PassengerId{},
	}
	err = pickup.Driver.WithSessionId(sessionId).Select()
	if err != nil {
		logger.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = GetJsonBody(request, pickup.PassengerId)
	if err != nil {
		logger.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := storage.NewUser().WithUserId(pickup.PassengerId.UUId)
	err = user.Select()
	if err != nil {
		logger.Warn(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	passenger := storage.NewPassenger().WithSessionId(user.SessionId).WithUserId(user.UserId)
	err = passenger.Delete()
	if err != nil {
		logger.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func GetPassengerInfo(writer http.ResponseWriter, request *http.Request) {
	sessionId, err := GetSessionId(request)
	if err != nil {
		logger.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	driver := storage.NewDriver()
	err = driver.WithSessionId(sessionId).Select()
	if err != nil {
		logger.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var mappings []*storage.Mapping
	err = storage.SelectDriverMapping(driver.UserId, &mappings)
	if err != nil {
		logger.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	var information []*storage.DriverInfo
	for _, mapping := range mappings {
		user := storage.NewUser()
		err = user.WithUserId(mapping.PassengerId.UUId).Select()
		if err != nil {
			logger.Error(err)
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
		logger.Error(err)
		WriteHttpResponse(writer, BadRequest)
		return
	}
	driver := storage.NewDriver()
	err = driver.WithSessionId(sessionId).Select()
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	var locations []storage.LocationLongLat
	err = GetJsonBody(request, &locations)
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	err = driver.WithLocations(&locations).Update()
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	WriteHttpResponse(writer, StatusOk)
}

func UpdateEstimations(writer http.ResponseWriter, request *http.Request) {
	sessionId, err := GetSessionId(request)
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, BadRequest)
		return
	}
	driver := storage.NewDriver()
	err = driver.WithSessionId(sessionId).Select()
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	var estimations []*storage.Estimation
	err = GetJsonBody(request, &estimations)
	for _, estimation := range estimations {
		mapping := storage.NewMapping(driver.UserId, estimation.PassengerId.UUId)
		err = mapping.Select()
		if err != nil {
			logger.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = mapping.WithTimes(estimation.PickupTime, estimation.DestinationTime).Update()
		if err != nil {
			logger.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	WriteHttpResponse(writer, StatusOk)
}

func ConfirmRideRequest(writer http.ResponseWriter, request *http.Request) {
	sessionId, err := GetSessionId(request)
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, BadRequest)
		return
	}
	driver := storage.NewDriver()
	err = driver.WithSessionId(sessionId).Select()
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	var confirmations []*storage.Confirmation
	err = GetJsonBody(request, &confirmations)
	for _, confirmation := range confirmations {
		mapping := storage.NewMapping(driver.UserId, confirmation.PassengerId.UUId)
		err = mapping.Select()
		if err != nil {
			logger.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		if confirmation.Accepted {
			err = mapping.WithStatus(storage.Accepted).Update()
			if err != nil {
				logger.Error(err)
				WriteHttpResponse(writer, InternalServerError)
				return
			}
			user := storage.NewUser()
			err = user.WithUserId(confirmation.PassengerId.UUId).Select()
			if err != nil {
				logger.Error(err)
				WriteHttpResponse(writer, InternalServerError)
				return
			}
			passenger := storage.NewPassenger()
			err = passenger.WithSessionId(user.SessionId).Select()
			if err != nil {
				logger.Error(err)
				WriteHttpResponse(writer, InternalServerError)
				return
			}
			err = driver.WithSeats(driver.Seats - passenger.RequestedSeats).Update()
			if err != nil {
				logger.Error(err)
				WriteHttpResponse(writer, InternalServerError)
				return
			}
		} else {
			err = mapping.WithStatus(storage.Denied).Update()
		}
		if err != nil {
			logger.Error(err)
			WriteHttpResponse(writer, InternalServerError)
			return
		}
	}
	WriteHttpResponse(writer, StatusOk)
}
