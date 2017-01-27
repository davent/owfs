package main

import (
	"log"

	"github.com/davent/owfs"
)

func main() {

	// Get a list of 1-wire devices from OWFS server
	devices, err := owfs.Dir("/")
	if err != nil {
		log.Printf("Could not get list of devices: %s", err)
	}

	// list each device
	for _, device := range devices {
		device.Temperature()
		log.Printf("Device: %+v", device)
	}

}
