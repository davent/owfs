package owfs

import (
	"log"
	"strconv"
	"strings"
)

type Device interface {
	Temperature() (float64, error)
}

type DS18S20 struct {
	ID     string
	Path   string
	Family int
	Temp   float64
}

func (d *DS18S20) Temperature() (temp float64, err error) {

	// Get data from OWFS server
	path := strings.Trim(d.Path, "\n") + "/temperature"
	data, err := Get(path)
	if err != nil {
		return 0, err
	}

	temp, err = strconv.ParseFloat(data, 64)

	d.Temp = temp

	return
}

func NewDevice(path string) (device Device, err error) {

	path = strings.TrimRight(path, "\x00")

	// Get id from path
	sep := strings.Split(path, ".")

	// Get family ID and device ID
	family, err := strconv.Atoi(strings.TrimLeft(sep[0], "/"))
	if err != nil {
		return nil, err
	}
	id := sep[1]

	switch family {
	case 10:
		device = &DS18S20{
			ID:     id,
			Path:   path,
			Family: family,
		}
		device.Temperature()
		return
	default:
		log.Printf("Family ID %d not currently supported", family)
	}

	return
}
