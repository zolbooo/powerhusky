package core

import (
	"encoding/json"
	"os"
)

type Config struct {
	DisableAutoShutdown bool
}

func ParseConfig() (*Config, error) {
	file, err := os.Open("/etc/powerhusky.conf.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	if err = json.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}
