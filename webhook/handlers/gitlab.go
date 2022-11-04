package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/zolbooo/powerhusky/webhook/core"
)

func GitlabWebhookHandler(w http.ResponseWriter, r *http.Request) {
	hook, _ := gitlab.New(gitlab.Options.Secret(os.Getenv(core.GITLAB_TOKEN)))

	payload, err := hook.Parse(r, gitlab.JobEvents)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jobEvent := payload.(gitlab.JobEventPayload)
	log.Printf("Job %d is using runner %d, status is %s", jobEvent.BuildID, jobEvent.Runner.ID, jobEvent.BuildStatus)
}
