package owfs

import (
	"testing"
)

const (
	DEFAULT_HOST string = "127.0.0.1"
	OWFS_HOST    string = "10.0.1.12"
)

func TestConfig(t *testing.T) {

	config := DefaultOWFSConfig()
	if config.Host != DEFAULT_HOST {
		t.Errorf("Default Host should be %s but was: %s", DEFAULT_HOST, config.Host)
	}

	// Override the default config
	Config(config)

}
