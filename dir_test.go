package owfs

import (
	"log"
	"testing"
)

func TestDir(t *testing.T) {

	// Get a list of all 1-wire devices from OWFS server
	devices, err := Dir("/")
	if err != nil {
		t.Errorf("Could not get list of devices: %s", err)
	}

	for _, device := range devices {
		log.Printf("Device: %+v", device)
	}

}
