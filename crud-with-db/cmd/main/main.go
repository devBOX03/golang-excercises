package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devBOX003/golang-excercises/crud-with-db/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 8000 \n")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
