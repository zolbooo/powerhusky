package webhook

import (
	"context"
	"io"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/zolbooo/powerhusky/webhook/core"
)

func init() {
	functions.HTTP("powerhusky", rootHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/health" {
		core.healthHandler(w, r)
		return
	}

	if r.URL.Path == "/start" {
		err := core.StartInstance(context.TODO())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, err.Error())
		} else {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "OK")
		}
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
