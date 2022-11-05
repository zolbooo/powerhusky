package handlers

import (
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

func GitlabWebhookHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := hook.Parse(r, gitlab.JobEvents)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Failed to parse request payload: %v", err)
		return
	}

	jobEvent := payload.(gitlab.JobEventPayload)
	log.Printf("Job %d is using runner %d, status is %s", jobEvent.BuildID, jobEvent.Runner.ID, jobEvent.BuildStatus)
}
