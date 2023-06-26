package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello!")
}

func signup(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Fprintf(w, "Path: %v \n", path)
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
	}
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %s \n", name)
	fmt.Fprintf(w, "Address: %s \n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/signup", signup)
	fmt.Println("Starting server at port 8000 ...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
