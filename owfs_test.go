package owfs

import (
	"log"
	"testing"
)

func TestSend(t *testing.T) {

	// Set Payload to be string of the path we wish list
	payload := []byte("")

	// New request header
	requestHeader := &RequestHeader{
		Type:          TYPE_NOOP,
		PayloadLength: uint32(len(payload)),
	}

	// Send to OWFS server
	responses, _, err := Send(requestHeader, payload)
	if err != nil {
		t.Errorf("Could not send data to OWFS server: %s", err)
	}

	for _, response := range responses {
		if response.Ret != 0 {
			t.Errorf("The response returned a non-zero return code: %d", response.Ret)
		}

		log.Printf("Response received: %+v", response)
	}

}
