package handlers

import (
	"testing"

	"github.com/go-playground/webhooks/v6/gitlab"
)

func TestIgnoredGitlabEvents(t *testing.T) {
	if !shouldIgnoreEvent(gitlab.BuildEventPayload{BuildStatus: "created"}) {
		t.Error("created build events should be ignored")
	}
	if !shouldIgnoreEvent(gitlab.BuildEventPayload{BuildStatus: "running"}) {
		t.Error("running build events should be ignored")
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
