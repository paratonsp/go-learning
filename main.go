package main

import (
	"fmt"
	"log"
	"net/http"

	"learning-go/middlewares"

	"learning-go/controllers/authcontroller"
	"learning-go/controllers/productcontroller"
	"learning-go/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>learning-go</h1>")
	})

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Get).Methods("GET")
	api.HandleFunc("/products/{id}", productcontroller.GetById).Methods("GET")
	api.HandleFunc("/products", productcontroller.Post).Methods("POST")
	api.HandleFunc("/products/{id}", productcontroller.Update).Methods("PATCH")
	api.HandleFunc("/products/{id}", productcontroller.Delete).Methods("DELETE")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":4000", r))
}
