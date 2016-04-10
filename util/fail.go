package util

import (
	log "github.com/Sirupsen/logrus"
)

func PanicOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

