package main

import (
	"log"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(interface{})
}

func main() {
	// define routes here
	route("GET", "/status", status)

	log.Fatal(
		http.ListenAndServe("0.0.0.0:4444", nil))
}
