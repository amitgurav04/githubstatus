package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

const (
	GITHUB_URL = "https://kctbh9vrtdwd.statuspage.io/api/v2/status.json"
)

// getStatus Handler
func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := http.Get(GITHUB_URL)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	githubStatus, err := parseGithubStatus(body)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	health := checkHealth(githubStatus)

	json.NewEncoder(w).Encode(health)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/status", getStatus).Methods("GET")
	http.ListenAndServe(":8000", router)

}
