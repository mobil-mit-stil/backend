package main

import (
	"backend/handler"
	"backend/storage"
	"backend/storage/memory"
	"fmt"
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"path"
	"runtime"
)

func main() {
	logger.SetReportCaller(true)
	logger.SetFormatter(&logger.TextFormatter{
		DisableQuote:  true,
		FullTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, filename := path.Split(f.File)
			caller := fmt.Sprint(" ", filename, ":", f.Line, ":")
			return caller, ""
		},
	})

	err := storage.Init(memory.New())
	if err != nil {
		logger.WithField("error", err).Fatal("could not init storage")
	}

	router := mux.NewRouter()
	handler.Register(router)

	port := ":8080"
	logger.WithField("port", port).Info("starting server")
	logger.Fatal(http.ListenAndServe(port, router))
}
