package handler

import (
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

func GetSessionId(r *http.Request) (string, error) {
    sessionId := r.Header.Get(HeaderAuthorization)
    if _, err := uuid.Parse(sessionId); err != nil {
        return "", err
    }
    return sessionId, nil
}


func WriteJSON(w http.ResponseWriter, response interface{}) {
    w.Header().Set(HeaderContentType, ContentTypeJSON)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}
