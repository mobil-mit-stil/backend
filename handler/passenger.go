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
		WriteError(writer, BadRequest)
		return
	}
	err = userPassenger.User.WithSessionId(userPassenger.Passenger.Session.Id).Create()
	if err != nil {
		logger.Error(err)
		WriteError(writer, InternalServerError)
		return
	}
	err = userPassenger.Passenger.WithUserId(userPassenger.User.UserId).Create()
	if err != nil {
		logger.Error(err)
		WriteError(writer, InternalServerError)
		return
	}
	WriteJSON(writer, userPassenger.Passenger.Session)
}

func RequestRide(writer http.ResponseWriter, request *http.Request) {

}

func GetDriverInfo(writer http.ResponseWriter, request *http.Request) {

}

func UpdatePassengerLocation(writer http.ResponseWriter, request *http.Request) {

}
