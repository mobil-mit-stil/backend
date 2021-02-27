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

func StartDriverSession(writer http.ResponseWriter, request *http.Request) {
	userDriver := UserDriver{
		User:   storage.NewUser(),
		Driver: storage.NewDriver(),
	}
	err := GetJsonBody(request, userDriver)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
	}
	err = userDriver.User.WithSessionId(userDriver.Driver.Session.Id).Create()
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
	err = userDriver.Driver.WithUserId(userDriver.User.UserId).Create()
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
	WriteJSON(writer, userDriver.Driver.Session)
}

func ConfirmPickup(writer http.ResponseWriter, request *http.Request) {

}

func GetPassengerInfo(writer http.ResponseWriter, request *http.Request) {

}

func UpdateRouteLocations(writer http.ResponseWriter, request *http.Request) {

}

func UpdateEstimations(writer http.ResponseWriter, request *http.Request) {

}

func ConfirmRideRequest(writer http.ResponseWriter, request *http.Request) {

}
