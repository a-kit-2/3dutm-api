package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/getData", GetData).Methods("GET")
	r.HandleFunc("/registerData", RegisterData).Methods("POST")

	http.ListenAndServe(":8000", r)
}
