package contollers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/devBOX003/golang-excercises/crud-with-db/pkg/models"
	"github.com/devBOX003/golang-excercises/crud-with-db/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, err := json.Marshal(books)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		bookDetails := models.GetBookById(id)
		res, err := json.Marshal(bookDetails)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	if book.Author == "" || book.Name == "" || book.Publication == "" {
		fmt.Println("Missing fields")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bookResult := book.CreateBook()
	res, err := json.Marshal(bookResult)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook *models.Book = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bookDetails := models.GetBookById(id)
	if updateBook.Author == "" || updateBook.Name == "" || updateBook.Publication == "" {
		fmt.Println("Missing fields")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if bookDetails.Name != updateBook.Name {
		bookDetails.Name = updateBook.Name
	}
	if bookDetails.Author != updateBook.Author {
		bookDetails.Author = updateBook.Author
	}
	if bookDetails.Publication != updateBook.Publication {
		bookDetails.Publication = updateBook.Publication
	}

	updatedBook := bookDetails.UpdateBook()
	res, err := json.Marshal(updatedBook)
	if err != nil {
		fmt.Printf("Error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		bookDetails := models.DeleteBookById(id)
		res, err := json.Marshal(bookDetails)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
