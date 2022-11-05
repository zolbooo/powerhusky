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

func CancelShutdown() {
	if shutdownTask != nil {
		shutdownTask.Cancel()
		shutdownTask = nil
	}
}

func ScheduleShutdown() (*time.Time, error) {
	// Debounce
	CancelShutdown()

	shutdownTime := time.Now().Add(time.Hour)

	task, err := taskScheduler.Schedule(func(ctx context.Context) {
		Shutdown()
	}, chrono.WithTime(shutdownTime))
	if err != nil {
		return nil, err
	}
	shutdownTask = task

	return &shutdownTime, nil
}
