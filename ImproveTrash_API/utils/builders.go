package utils

import (
	"ImproveTrash_API/handlers"
	"github.com/gorilla/mux"
)

// Отдельные функции для роутинга

func BuildBookResource(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"/{id}", handlers.GetBookById).Methods("GET")
	router.HandleFunc(prefix, handlers.CreateBook).Methods("POST")
	router.HandleFunc(prefix+"/{id}", handlers.UpdateBookById).Methods("PUT")
	router.HandleFunc(prefix+"/{id}", handlers.DeleteBookById).Methods("DELETE")
}

func BuildManyBooksResourcePrefix(router *mux.Router, prefix string) {
	router.HandleFunc(prefix+"s", handlers.GetAllBooks).Methods("GET")
}