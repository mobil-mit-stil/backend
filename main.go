package main

import (
    "backend/handler"
    "backend/storage"
    "backend/storage/memory"
    "fmt"
    logger "github.com/sirupsen/logrus"
    "github.com/gorilla/mux"
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
    logger.Fatal(http.ListenAndServe(":8080", router))
}
