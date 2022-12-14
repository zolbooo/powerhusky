package core

import (
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	Rpc struct {
		Port  int
		Token string
		Tls   struct {
			CertFile string
			KeyFile  string
		}
	}

	DisableAutoShutdown bool
	CounterPath         string
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

	if config.Rpc.Port == 0 {
		config.Rpc.Port = 2333
	}
	if config.Rpc.Token == "" {
		return nil, errors.New("no RPC token provided")
	}
	if config.CounterPath == "" {
		config.CounterPath = "/var/run/powerhusky.pid"
	}
	return config, nil
}
