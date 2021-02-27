package handler

import (
	"backend/storage"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(baseRouter *mux.Router) {
	d := baseRouter.
		PathPrefix("/driver").
		Subrouter()
	{
		d.
			Path("/start").
			Methods(http.MethodPost).
			HandlerFunc(StartDriverSession)
		d.
			Path("/pickup").
			Methods(http.MethodPost).
			HandlerFunc(ConfirmPickup)
		d.
			Path("/information").
			Methods(http.MethodGet).
			HandlerFunc(GetPassengerInfo)
		d.
			Path("/locations").
			Methods(http.MethodPost).
			HandlerFunc(UpdateRouteLocations)
		d.
			Path("/estimations").
			Methods(http.MethodPost).
			HandlerFunc(UpdateEstimations)
		d.
			Path("/confirmations").
			Methods(http.MethodPost).
			HandlerFunc(ConfirmRideRequest)
	}

	p := baseRouter.
		PathPrefix("/passenger").
		Subrouter()
	{
		p.
			Path("/start").
			Methods(http.MethodPost).
			HandlerFunc(StartPassengerSession)
		p.
			Path("/request").
			Methods(http.MethodPost).
			HandlerFunc(RequestRide)
		p.
			Path("/information").
			Methods(http.MethodGet).
			HandlerFunc(GetDriverInfo)
		p.
			Path("/location").
			Methods(http.MethodPost).
			HandlerFunc(UpdatePassengerLocation)
	}

	baseRouter.
		Path("/debug").
		Methods(http.MethodGet).
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			var users []*storage.User
			_ = storage.GetProvider().SelectUsers(&users)
			WriteJSON(writer, users)

			var drivers []*storage.Driver
			_ = storage.GetProvider().SelectDrivers(&drivers)
			WriteJSON(writer, drivers)

			var passengers []*storage.Passenger
			_ = storage.GetProvider().SelectPassengers(&passengers)
			WriteJSON(writer, passengers)

			var mappings []*storage.Mapping
			_ = storage.GetProvider().SelectMappings(&mappings)
			WriteJSON(writer, mappings)

		})
}
