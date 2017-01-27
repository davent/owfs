package owfs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type Client struct {
	Config *OWFSConfig
}

var owfs_client *Client

func init() {
	owfs_client = &Client{
		Config: DefaultOWFSConfig(),
	}
}

func (c *Client) getConn() (conn net.Conn, err error) {

	conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", c.Config.Host, c.Config.Port))
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to OWFS server: %s", err)
	}

	return conn, nil

}

func Send(request_header *RequestHeader, payload []byte) (response_headers []*ResponseHeader, response_payloads [][]byte, err error) {
	response_headers, response_payloads, err = owfs_client.send(request_header, payload)
	return
}
func (c *Client) send(request_header *RequestHeader, payload []byte) (response_headers []*ResponseHeader, response_payloads [][]byte, err error) {

	// Get a new connection to the OWFS server
	conn, err := c.getConn()
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()

	buf := new(bytes.Buffer)
	// Send the request header to the OWFS server
	err = binary.Write(buf, binary.BigEndian, request_header)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to send header to OWFS server: %s", err)
	}

	// Send the request payload to the OWFS server
	err = binary.Write(buf, binary.BigEndian, payload)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to send payload to OWFS server: %s", err)
	}
	conn.Write(buf.Bytes())

	// Get reponses
	for {

		// Create a new reponse header
		response_header := &ResponseHeader{
			Ret: 1, // Default to erroneous state
		}

		// Read the response header
		err = binary.Read(conn, binary.BigEndian, response_header)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, nil, fmt.Errorf("Failed to get response from OWFS server: %s", err)
		}

		if response_header.PayloadLength > 0 {

			// Read the payload
			buf := make([]byte, response_header.PayloadLength)
			err = binary.Read(conn, binary.BigEndian, &buf)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return nil, nil, fmt.Errorf("Failed to read payload from OWFS server: %s", err)
				}
			}

			response_headers = append(response_headers, response_header)
			response_payloads = append(response_payloads, buf)

		} else {
			break
		}
	}
	err = nil
	return
}
