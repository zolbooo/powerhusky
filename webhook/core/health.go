package core

import (
	"io"
	"log"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := ConfigFromEnv(); err != nil {
		w.WriteHeader(http.StatusConflict)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}
