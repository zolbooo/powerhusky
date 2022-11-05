package core

import (
	"context"
	"time"

	"github.com/procyon-projects/chrono"
)

var (
	taskScheduler                      = chrono.NewDefaultTaskScheduler()
	shutdownTask  chrono.ScheduledTask = nil
)

func ScheduleShutdown() error {
	if shutdownTask != nil {
		// Debounce
		shutdownTask.Cancel()
	}

	// Schedule shutdown after 1 hour by default
	task, err := taskScheduler.Schedule(func(ctx context.Context) {
		Shutdown()
	}, chrono.WithTime(time.Now().Add(time.Hour)))
	if err != nil {
		return err
	}
	shutdownTask = task
	return nil
}
