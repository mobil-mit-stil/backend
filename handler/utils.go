package handler

import (
	"backend/storage"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	HeaderAuthorization = "Authorization"
	HeaderContentType   = "Content-Type"
	ContentTypeJSON     = "application/json"
)

func GetJsonBody(r *http.Request, ref interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ref)
	if err != nil {
		logrus.Error(err)
		return fmt.Errorf("invalid json data")
	}
	return nil
}

func GetSessionId(r *http.Request) (storage.SessionUUId, error) {
	sessionId := r.Header.Get(HeaderAuthorization)
	if _, err := uuid.Parse(sessionId); err != nil {
		return "", err
	}
	return storage.SessionUUId(sessionId), nil
}

func WriteJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set(HeaderContentType, ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func WriteHttpResponse(w http.ResponseWriter, res HttpResponse) {
	w.Header().Set(HeaderContentType, ContentTypeJSON)
	w.WriteHeader(res.StatusCode)
	json.NewEncoder(w).Encode(res)
}

type HttpResponse struct {
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"msg"`
}

var StatusOk = HttpResponse{
	StatusCode: 200,
	Msg:        "Ok",
}

var BadRequest = HttpResponse{
	StatusCode: 400,
	Msg:        "Bad Request",
}

var InternalServerError = HttpResponse{
	StatusCode: 500,
	Msg:        "Internal Server Error",
}
