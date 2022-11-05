package daemon

import (
	"context"

	"github.com/kardianos/service"
)

var ServiceConfig = &service.Config{
	Name:        "powerhusky",
	DisplayName: "Powerhusky",
	Description: "Power management daemon",
}

type Service struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (s *Service) Start(svg service.Service) error {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	// Start should not block. Do the actual work async.
	go func() {
		// TODO: Implement service
	}()
	return nil
}
func (s *Service) Stop(svc service.Service) error {
	s.cancel()
	return nil
}
