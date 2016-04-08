package util

import (
	"log"
	"fmt"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func Fail(msg string) {
	log.Fatalf("%s", msg)
	panic(fmt.Sprintf("%s", msg))
}
