package webhook

import (
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("powerhusky", rootHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/health" {
		healthHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
