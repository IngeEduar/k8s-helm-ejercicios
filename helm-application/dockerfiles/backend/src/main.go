package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type HandsOn struct {
	Time     time.Time `json:"time"`
	Hostname string    `json:"hostname"`
}

func ServerHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp := HandsOn{
		Time:     time.Now(),
		Hostname: os.Getenv("HOSTNAME"),
	}
	jsonResp, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET") // Puedes ajustar según tus necesidades
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonResp))
}

func main() {
	http.HandleFunc("/", ServerHTTP)
	http.ListenAndServe(":8080", nil)
}
