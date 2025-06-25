package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	currentTime := time.Now().Format(time.RFC3339)
	response := TimeResponse{Time: currentTime}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/time", timeHandler)
	log.Fatal(http.ListenAndServe(":8795", nil))
}
