package owfs

import (
	"log"
	"testing"
)

func TestGet(t *testing.T) {

	// Get a list of all 1-wire devices from OWFS server
	response, err := Get("/10.67C6697351FF/temperature")
	if err != nil {
		t.Errorf("Could not get device data: %s", err)
	}

	log.Printf("Response: %+v", response)

}
