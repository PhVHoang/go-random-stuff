package main

import (
	"os"
	"runtime/debug"

	log "github.com/sirupsen/logrus"
)

func PanicHandler() {
	r := recover()

	if r == nil {
		return // no pic underway
	}

	log.Errorf("Panic occurred in this application %v\n", r)

	// print debug stack
	debug.PrintStack()

	os.Exit(1)
}
