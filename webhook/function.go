package webhook

import (
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/zolbooo/powerhusky/webhook/core"
	"github.com/zolbooo/powerhusky/webhook/handlers"
)

func init() {
	functions.HTTP("powerhusky", rootHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/health" {
		core.HealthHandler(w, r)
		return
	}

	if r.URL.Path == "/webhook/gitlab" {
		handlers.GitlabWebhookHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
