package daemon

import (
	"context"
	"os"

	"github.com/kardianos/service"
	"github.com/zolbooo/powerhusky/daemon/core"
	"github.com/zolbooo/powerhusky/daemon/rpc"
)

var ServiceConfig = &service.Config{
	Name:        "powerhusky",
	DisplayName: "Powerhusky",
	Description: "Power management daemon",
}

type Service struct {
	ctx    context.Context
	cancel context.CancelFunc

	Logger service.Logger
}

func (s *Service) run() {
	config, err := core.ParseConfig()
	if err != nil {
		s.Logger.Errorf("failed to parse config: %v", err)
	}
	if config == nil {
		s.Logger.Error("Fatal: no required config options are provided")
		return
	}

	if !config.DisableAutoShutdown {
		shutdownTime, err := core.ScheduleShutdown()
		if err != nil {
			s.Logger.Errorf("failed to schedule shutdown: %v", err)
		} else {
			s.Logger.Infof("Shutdown was scheduled to %s", shutdownTime.String())
		}
	}

	counterData := &core.CounterData{Counter: 1, Pid: os.Getpid()}
	if err = counterData.Save(config.CounterPath); err != nil {
		s.Logger.Errorf("Fatal: failed to save counter file: %v", err)
		return
	}

	var tlsOptions *rpc.TLSOptions = nil
	if config.Rpc.Tls.CertFile != "" && config.Rpc.Tls.KeyFile != "" {
		tlsOptions = &rpc.TLSOptions{CertFile: tlsOptions.CertFile, KeyFile: config.Rpc.Tls.KeyFile}
	}
	rpc.InitServer(s.ctx, config.Rpc.Token, config.Rpc.Port, tlsOptions)
	s.Logger.Infof("RPC running on port %d, using TLS: %v", config.Rpc.Port, tlsOptions != nil)
}

func (s *Service) Start(svc service.Service) error {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	// Start should not block. Do the actual work async.
	go s.run()
	return nil
}
func (s *Service) Stop(svc service.Service) error {
	s.cancel()
	return nil
}
