package main

import (
	"log"

	"github.com/davent/owfs"
)

func main() {

	// The ID of the 1-wire device
	device := "28.EEF5B31B1601DA"

	// Use a remote OWFS server instead of the default localhost
	config := owfs.DefaultOWFSConfig()
	config.Host = "10.0.1.12"

	// Override the default config
	owfs.Config(config)

	// Get the temp from a known device
	temp, err := owfs.Get("/" + device + "/temperature")
	if err != nil {
		log.Printf("Could not get temperatures from %s: %s", device, err)
	}

	// Print result
	log.Printf("Temperature: %s°C", temp)

}
