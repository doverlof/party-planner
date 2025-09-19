package main

import (
	""
	"github.com/doverlof//books"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	database := db.Connect("mongodb://localhost:27017")

	repo := books.NewRepository(database)
	handler := books.NewHandler(repo)

	r := mux.NewRouter()
	r.HandleFunc("/books", handler.GetBooks).Methods("GET")
	r.HandleFunc("/books", handler.CreateBook).Methods("POST")

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
