package utils

import "log"

var logFatalFunc = log.Fatal

func HandleError(err error) {
	if err != nil {
		logFatalFunc("erro: ", err)
	}
}
