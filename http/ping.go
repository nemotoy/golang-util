package ping

import (
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("pong"))
}
