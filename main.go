package main

import (
    "fmt"
    logger "github.com/sirupsen/logrus"
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
}
