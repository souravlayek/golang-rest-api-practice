package router

import (
	"github.com/gorilla/mux"
	"github.com/souravlayek/rest-api-tutorial/internal/api/handler"
)

func CreateBookRoutes(router *mux.Router) {
	router.HandleFunc("/api/books", handler.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", handler.GetOneBook).Methods("GET")
	router.HandleFunc("/api/books", handler.InsertBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", handler.UpdateOneBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", handler.DeleteOneBook).Methods("DELETE")
}
