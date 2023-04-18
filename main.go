package main

import (
	"fmt"

	"learning-go/controllers"
	"learning-go/middlewares"
	"learning-go/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "learning-go")
	})

	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", controllers.GetProduct).Methods("GET")
	api.HandleFunc("/products/{id}", controllers.GetProductById).Methods("GET")
	api.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	api.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PATCH")
	api.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
