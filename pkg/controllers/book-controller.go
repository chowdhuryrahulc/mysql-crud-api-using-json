package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/chowdhuryrahulc/mysql-crud-api-using-json/pkg/utils"
	"github.com/chowdhuryrahulc/mysql-crud-api-using-json/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)	// Marshel converts go object into json
	w.Header().Set("Content-Type", "publication/json")
	w.WriteHeader(http.StatusOK)		// means 200
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)					// mux.Vars(r): extracts all the key value pairs from the requests
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "publication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	// we recieve json, we changed (parsed) it to something db will understand (go-object), and send to db. Db send the same thing back.
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)		// we need to Marshel something before sending back to user. And unmarshel something after recieving something
	// w.Header().Set("Content-Type", "application/json")	//! might be wrong, delete later
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)			// book is the book that has been deleted. Now we send it back to user
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "publication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	// mix of CreateBook() and GetBookById() functions
	// we search book by id, and update its values. And then save it. Then just send the response back to user.
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)					// mux.Vars(r): extracts all the key value pairs from the requests
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	booksDetails, db := models.GetBookById(ID)
	if updateBook.Name != ""{
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		booksDetails.Publication = updateBook.Publication
	}
	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type", "publication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}