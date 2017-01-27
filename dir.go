package owfs

func Dir(path string) (devices []Device, err error) {
	devices, err = owfs_client.dir(path)
	return
}
func (c *Client) dir(path string) (devices []Device, err error) {

	// Set payload as path to dir
	payload := []byte(path)

	// New request header
	request_header := &RequestHeader{
		Type:          TYPE_DIR,
		PayloadLength: uint32(len(payload)),
	}

	// Send request header and payload to OWFS server
	response_headers, response_payloads, err := c.send(request_header, payload)
	if err != nil {
		return nil, err
	}

	for i, response_header := range response_headers {
		if response_header.Ret == 0 {
			device, err := NewDevice(string(response_payloads[i]))
			if err != nil {
				return nil, err
			}

			devices = append(devices, device)
		}
	}

	return
}
