package webhook

import (
	"io"
	"net/http"
	"os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("GITLAB_TOKEN") == "" {
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, "GITLAB_TOKEN environment is not defined")
		return
	}
	if os.Getenv("GCP_API_TOKEN") == "" {
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, "GCP_API_TOKEN environment is not defined")
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}
