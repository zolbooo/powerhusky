package core

import (
	"errors"
	"os"
)

const (
	GITLAB_TOKEN  = "GITLAB_TOKEN"
	DAEMON_SECRET = "DAEMON_SECRET"

	GCP_PROJECT         = "GCP_PROJECT"
	GCE_INSTANCE_ID     = "GCE_INSTANCE_ID"
	GCE_INSTANCE_REGION = "GCE_INSTANCE_REGION"
)

type Config struct {
	GitlabToken  string
	DaemonSecret string

	GCPProject        string
	GCEInstanceID     string
	GCEInstanceRegion string
}

func ConfigFromEnv() (*Config, error) {
	config := &Config{
		GitlabToken:       os.Getenv(GITLAB_TOKEN),
		DaemonSecret:      os.Getenv(DAEMON_SECRET),
		GCPProject:        os.Getenv(GCP_PROJECT),
		GCEInstanceID:     os.Getenv(GCE_INSTANCE_ID),
		GCEInstanceRegion: os.Getenv(GCE_INSTANCE_REGION),
	}
	if config.GitlabToken == "" {
		return nil, errors.New("GITLAB_TOKEN environment variable is not defined")
	}
	if config.DaemonSecret == "" {
		return nil, errors.New("DAEMON_SECRET environment variable is not defined")
	}
	if config.GCPProject == "" {
		return nil, errors.New("GCP_PROJECT environment variable is not defined")
	}
	if config.GCEInstanceID == "" {
		return nil, errors.New("GCP_INSTANCE_ID environment variable is not defined")
	}
	if config.GCEInstanceRegion == "" {
		return nil, errors.New("GCP_INSTANCE_REGION environment variable is not defined")
	}
	return config, nil
}

var appConfig *Config

func init() {
	appConfig, _ = ConfigFromEnv()
}
