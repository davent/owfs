package owfs

import (
	"strings"
)

func Get(path string) (response string, err error) {
	response, err = owfs_client.get(path)
	return
}
func (c *Client) get(path string) (response string, err error) {

	// Set payload as path to device
	payload := make([]byte, len(path)+1) // To get null terminated
	copy(payload, []byte(path))

	// New request header
	request_header := &RequestHeader{
		Type:          TYPE_GET,
		PayloadLength: uint32(len(payload)),
		Size:          8192,
	}

	// Send request header and payload to OWFS server
	_, response_payloads, err := c.send(request_header, payload)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(response_payloads[0])), nil

}
