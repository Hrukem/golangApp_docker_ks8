package handlers

import "net/http"

// healthz is a liveliness probe
func healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
