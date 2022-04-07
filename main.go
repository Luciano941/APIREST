package main

import (
	"fmt"
	"log"
	"net/http"
	"packages/handlers"
	"packages/models"

	"github.com/gorilla/mux"
)

func main() {

	models.Migrate()

	mux := mux.NewRouter()

	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	fmt.Println("Servidor corriendo en localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
