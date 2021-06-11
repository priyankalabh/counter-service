package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
{
	"count": 1
}
*/
var counter int = 0

func main() {
	http.HandleFunc("/api/count", countHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Listen and serve error  %v", err)

	}
}

type countResponse struct {
	Count int `json:"count"`
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	count := countResponse{
		Count: counter,
	}
	counter++
	num, err := json.Marshal(&count)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(num)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}
