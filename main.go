package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}", StoreBTC).Methods("POST")

	http.ListenAndServe(":80", r)
}

func StoreBTC(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// title := vars["title"]
	// page := vars["page"]

	req := requestBody{}
	json.Unmarshal(r.Body.Read(), req)
	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

type requestBody struct {
	DateTime time.Time `json:"datetime"`
	Amount   float64   `json:"amount"`
}
