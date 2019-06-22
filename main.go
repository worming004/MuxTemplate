package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Dummy", dummyHandler).Methods("GET")
	http.Handle("/", r)

	cors.New(cors.Options{AllowedOrigins: []string{"http://localhost:3000"}})

	http.ListenAndServe(":3000", r)

}

func dummyHandler(w http.ResponseWriter, r *http.Request) {
	dummy := buildDummy
	js, err := json.Marshal(dummy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type dummy struct {
	id   int
	name string
}

func buildDummy() dummy {
	return dummy{1, "DummyName"}
}
