package api

import (
	"GravitumTest/internal/api/handlers"
	"GravitumTest/internal/api/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartRouter(port string) {
	r := mux.NewRouter()
	r.Use(middleware.BasicAuth)
	r.Handle("/users/{id}", http.HandlerFunc(handlers.GetUser)).Methods("GET")
	r.Handle("/users", http.HandlerFunc(handlers.NewUser)).Methods("POST")
	r.Handle("/users/{id}", http.HandlerFunc(handlers.UpdateUser)).Methods("PUT")
	http.Handle("/", r)
	log.Printf("Starting server on %s \n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
