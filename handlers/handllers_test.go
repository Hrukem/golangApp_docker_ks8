package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRoute(t *testing.T) {
	r := Route("", "", "")
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/home")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf(`Status code for "/home" is wrong. Have: %d, want: %d`, res.StatusCode, http.StatusOK)
	}

	res, err = http.Post(ts.URL+"/home", "text/plain", nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf(`Status code for "/home" is wrong. Have: %d, want: %d`, res.StatusCode, http.StatusMethodNotAllowed)
	}

	res, err = http.Get(ts.URL + "/not-exists")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Errorf(`Status code for "/home" is wrong. Have: %d, want: %d`, res.StatusCode, http.StatusNotFound)
	}

}

func Test_home(t *testing.T) {
	w := httptest.NewRecorder()
	buildTime := time.Now().Format("20060102_03:04:05")
	commit := "some test hash"
	release := "0.0.8"

	h := home(buildTime, commit, release)
	h(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}

	greeting, err := io.ReadAll(resp.Body)
	if err = resp.Body.Close(); err != nil {
		log.Println("Error close Body in tests")
	}
	if err != nil {
		t.Fatal(err)
	}

	info := struct {
		BuildTime string `json:"buildTime"`
		Commit    string `json:"commit"`
		Release   string `json:"release"`
	}{}

	if err = json.Unmarshal(greeting, &info); err != nil {
		t.Fatal(err)
	}

	if info.Release != release {
		t.Errorf("Release version is wrong. Have: %s, want: %s", info.Release, release)
	}
	if info.BuildTime != buildTime {
		t.Errorf("Build time is wrong. Have: %s, want: %s", info.BuildTime, buildTime)
	}
	if info.Commit != commit {
		t.Errorf("Commit is wrong. Have: %s, want: %s", info.Commit, commit)
	}
}
