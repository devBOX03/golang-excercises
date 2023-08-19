package routes

import (
	"github.com/devBOX003/golang-excercises/crud-with-db/pkg/contollers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", contollers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", contollers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", contollers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", contollers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", contollers.DeleteBook).Methods("DELETE")
}
