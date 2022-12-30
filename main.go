package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get/{key}", GetData).Methods("GET")
	r.HandleFunc("/post", RegisterData).Methods("POST")

	c := cors.Default().Handler(r)

	http.ListenAndServe(":8000", c)
}
