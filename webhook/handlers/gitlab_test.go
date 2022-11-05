package handlers

import (
	"testing"

	"github.com/go-playground/webhooks/v6/gitlab"
)

func TestIgnoredGitlabEvents(t *testing.T) {
	if !shouldIgnoreEvent(gitlab.BuildEventPayload{BuildStatus: "created"}) {
		t.Error("created build events should be ignored")
	}

	if shouldIgnoreEvent(gitlab.BuildEventPayload{BuildStatus: "running"}) {
		t.Error("running build events should not be ignored")
	}
	if shouldIgnoreEvent(gitlab.BuildEventPayload{BuildStatus: "pending"}) {
		t.Error("pending build events should not be ignored")
	}
	if shouldIgnoreEvent(gitlab.BuildEventPayload{BuildStatus: "success"}) {
		t.Error("failed build events should not be ignored")
	}
	if shouldIgnoreEvent(gitlab.BuildEventPayload{BuildStatus: "failed"}) {
		t.Error("failed build events should not be ignored")
	}
}

func TestStartInstance(t *testing.T) {
	if shouldStartInstance(gitlab.BuildEventPayload{BuildStatus: "pending", Runner: gitlab.Runner{IsShared: true}}) {
		t.Error("instance must not be started when instance is shared")
	}
	if shouldStartInstance(gitlab.BuildEventPayload{BuildStatus: "running", Runner: gitlab.Runner{IsShared: true}}) {
		t.Error("instance must not be started when instance is shared and job is running")
	}
	if shouldStartInstance(gitlab.BuildEventPayload{BuildStatus: "failed", Runner: gitlab.Runner{IsShared: false}}) {
		t.Error("instance must not be started when job is failed")
	}

	if !shouldStartInstance(gitlab.BuildEventPayload{BuildStatus: "pending", Runner: gitlab.Runner{IsShared: false}}) {
		t.Error("instance must be started when instance is not shared")
	}
	if !shouldStartInstance(gitlab.BuildEventPayload{BuildStatus: "pending", Runner: gitlab.Runner{IsShared: false, ID: 1}}) {
		t.Error("instance must be started when instance is not shared and ID is defined")
	}
	if !shouldStartInstance(gitlab.BuildEventPayload{BuildStatus: "running", Runner: gitlab.Runner{IsShared: false, ID: 1}}) {
		t.Error("instance must be started when job is running on non-shared instance")
	}
	if !shouldStartInstance(gitlab.BuildEventPayload{BuildStatus: "running", Runner: gitlab.Runner{IsShared: false}}) {
		t.Error("instance must be started when job is running on non-shared instance")
	}
}
