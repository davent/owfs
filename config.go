package owfs

type OWFSConfig struct {
	Host string
	Port int
}

func DefaultOWFSConfig() (config *OWFSConfig) {

	config = &OWFSConfig{
		Host: "127.0.0.1",
		Port: 4304,
	}

	return
}

func Config(config *OWFSConfig) (err error) {
	owfs_client.Config = config

	return
}
