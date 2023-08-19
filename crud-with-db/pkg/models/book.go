package models

import (
	"github.com/devBOX003/golang-excercises/crud-with-db/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func (book *Book) UpdateBook() *Book {
	db.Save(&book)
	return book
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) *Book {
	var book Book
	db.Where("ID=?", id).Find(&book)
	return &book
}

func DeleteBookById(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
