package daemon

import (
	"context"

	"github.com/kardianos/service"
	"github.com/zolbooo/powerhusky/daemon/core"
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

func (s *Service) Start(svg service.Service) error {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	// Start should not block. Do the actual work async.
	go func() {
		config, err := core.ParseConfig()
		if err != nil {
			s.Logger.Errorf("failed to parse config: %v", err)
		}

		if !config.DisableAutoShutdown {
			if err = core.ScheduleShutdown(); err != nil {
				s.Logger.Errorf("failed to schedule shutdown: %v", err)
			}
		}
		// TODO: Implement RPC listener
	}()
	return nil
}
func (s *Service) Stop(svc service.Service) error {
	s.cancel()
	return nil
}
