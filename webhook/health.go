package webhook

import (
	"io"
	"net/http"
	"os"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("GITLAB_TOKEN") == "" {
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, "GITLAB_TOKEN environment variable is not defined")
		return
	}

	if os.Getenv("GCP_PROJECT") == "" {
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, "GCP_PROJECT environment variable is not defined. Are you running in the GCF environment?")
		return
	}
	if os.Getenv("GCE_INSTANCE_ID") == "" {
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, "GCP_INSTANCE_ID environment variable is not defined")
		return
	}
	if os.Getenv("GCE_INSTANCE_REGION") == "" {
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, "GCP_INSTANCE_REGION environment variable is not defined")
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}
