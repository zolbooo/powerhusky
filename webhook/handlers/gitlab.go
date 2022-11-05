package handlers

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/zolbooo/powerhusky/webhook/core"
)

var hook *gitlab.Webhook

func init() {
	hook, _ = gitlab.New(gitlab.Options.Secret(os.Getenv(core.GITLAB_TOKEN)))
}

func shouldIgnoreEvent(buildEvent gitlab.BuildEventPayload) bool {
	return buildEvent.BuildStatus == "created" || buildEvent.BuildStatus == "running"
}

func handleBuildEvent(ctx context.Context, buildEvent gitlab.BuildEventPayload) error {
	log.Printf("Job %d is using runner %d, status is %s", buildEvent.BuildID, buildEvent.Runner.ID, buildEvent.BuildStatus)
	if shouldIgnoreEvent(buildEvent) {
		return nil
	}
	if buildEvent.BuildStatus == "pending" {
		if buildEvent.Runner.ID == 0 && !buildEvent.Runner.IsShared {
			return core.StartInstance(ctx)
		}
	} else if !buildEvent.BuildFinishedAt.IsZero() && !buildEvent.Runner.IsShared {
		// Job has finished, stop instance
		return core.StopInstance(ctx)
	}
	return nil
}

func GitlabWebhookHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := hook.Parse(r, gitlab.JobEvents, gitlab.BuildEvents)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed to parse request payload: %v", err)
		return
	}
	if jobEvent, ok := payload.(gitlab.JobEventPayload); ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Unexpected job event: %+v", jobEvent)
		return
	}

	buildEvent, ok := payload.(gitlab.BuildEventPayload)
	if ok {
		err := handleBuildEvent(r.Context(), buildEvent)
		if err == context.DeadlineExceeded || err == context.Canceled {
			w.WriteHeader(http.StatusRequestTimeout)
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("failed to handle job event: %v", err)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}
