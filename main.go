package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paratonsp/learning-go/controllers/authcontroller"
	"github.com/paratonsp/learning-go/controllers/productcontroller"
	"github.com/paratonsp/learning-go/middlewares"
	"github.com/paratonsp/learning-go/models"
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

	log.Fatal(http.ListenAndServe(":8000", r))
}
