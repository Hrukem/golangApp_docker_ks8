package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func home(buildTIme, commit, release string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		info := struct {
			BuildTime string `json:"buildTime"`
			Commit    string `json:"commit"`
			Release   string `json:"release"`
		}{
			buildTIme, commit, release,
		}

		body, err := json.Marshal(info)
		if err != nil {
			log.Printf(`Error encode "info" in func "home": %v`, err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err = w.Write(body); err != nil {
			log.Printf("Error send http: %v", err)
		}
	}
}
