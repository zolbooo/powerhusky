package daemon

import (
	"context"

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

func (s *Service) Start(svc service.Service) error {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	// Start should not block. Do the actual work async.
	go func() {
		config, err := core.ParseConfig()
		if err != nil {
			s.Logger.Errorf("failed to parse config: %v", err)
		}
		if config == nil {
			s.Logger.Error("no required config options are provided, exiting")
			return
		}

		if !config.DisableAutoShutdown {
			if err = core.ScheduleShutdown(); err != nil {
				s.Logger.Errorf("failed to schedule shutdown: %v", err)
			}
		}

		var tlsOptions *rpc.TLSOptions = nil
		if config.Rpc.Tls.CertFile != "" && config.Rpc.Tls.KeyFile != "" {
			tlsOptions = &rpc.TLSOptions{CertFile: tlsOptions.CertFile, KeyFile: config.Rpc.Tls.KeyFile}
		}
		rpc.InitServer(s.ctx, config.Rpc.Token, config.Rpc.Port, tlsOptions)
		s.Logger.Infof("RPC running on port %d, using TLS: %v", config.Rpc.Port, tlsOptions != nil)
	}()
	return nil
}
func (s *Service) Stop(svc service.Service) error {
	s.cancel()
	return nil
}
