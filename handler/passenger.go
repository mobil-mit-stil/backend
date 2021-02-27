package handler

import (
	"backend/storage"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

type UserPassenger struct {
	*storage.User
	*storage.Passenger
}

func StartPassengerSession(writer http.ResponseWriter, request *http.Request) {
	userPassenger := &UserPassenger{
		User:      storage.NewUser(),
		Passenger: storage.NewPassenger(),
	}
	err := GetJsonBody(request, userPassenger)
	if err != nil {
		logger.Warn(err)
		WriteHttpResponse(writer, BadRequest)
		return
	}
	err = userPassenger.User.WithSessionId(userPassenger.Passenger.Session.Id).Create()
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	err = userPassenger.Passenger.WithUserId(userPassenger.User.UserId).Create()
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	WriteJSON(writer, userPassenger.Passenger.Session)
}

func RequestRide(writer http.ResponseWriter, request *http.Request) {
	sessionId, err := GetSessionId(request)
	if err != nil {
		logger.Warn(err)
		WriteHttpResponse(writer, BadRequest)
		return
	}
	passenger := storage.NewPassenger().WithSessionId(sessionId)
	err = passenger.Select()
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	driverId := &storage.DriverId{}
	err = GetJsonBody(request, driverId)
	if err != nil {
		logger.Warn(err)
		WriteHttpResponse(writer, BadRequest)
		return
	}
	mapping := storage.NewMapping(driverId.UUId, passenger.UserId)
	err = mapping.Select()
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	err = mapping.WithRequested(true).Update()
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	WriteHttpResponse(writer, StatusOk)
}

func GetDriverInfo(writer http.ResponseWriter, request *http.Request) {
	sessionId, err := GetSessionId(request)
	if err != nil {
		logger.Warn(err)
		WriteHttpResponse(writer, BadRequest)
		return
	}
	passenger := storage.NewPassenger().WithSessionId(sessionId)
	err = passenger.Select()
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	var mappings []*storage.Mapping
	err = storage.SelectPassengerMapping(passenger.UserId, &mappings)
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	var information []*storage.PassengerInfo
	for _, mapping := range mappings {
		user := storage.NewUser().WithUserId(mapping.DriverId.UUId)
		err = user.Select()
		if err != nil {
			logger.Error(err)
			WriteHttpResponse(writer, InternalServerError)
			return
		}
		information = append(information, &storage.PassengerInfo{
			DriverId:        mapping.DriverId,
			Name:            user.Name,
			PickupTime:      mapping.PickupTime,
			DestinationTime: mapping.DestinationTime,
		})
	}
	WriteJSON(writer, information)
}

func UpdatePassengerLocation(writer http.ResponseWriter, request *http.Request) {
	sessionId, err := GetSessionId(request)
	if err != nil {
		logger.Warn(err)
		WriteHttpResponse(writer, BadRequest)
		return
	}
	passenger := storage.NewPassenger().WithSessionId(sessionId)
	err = passenger.Select()
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	var location storage.LocationLongLat
	err = GetJsonBody(request, &location)
	if err != nil {
		logger.Warn(err)
		WriteHttpResponse(writer, BadRequest)
		return
	}
	err = passenger.WithLocation(&location).Update()
	if err != nil {
		logger.Error(err)
		WriteHttpResponse(writer, InternalServerError)
		return
	}
	WriteHttpResponse(writer, StatusOk)
}
