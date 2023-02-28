package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gitlab.com/nisarg/3/pkg/models"
	"gitlab.com/nisarg/3/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	NewBooks := models.GetAllBooks()
	res, _ := json.Marshal(NewBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	var ID string
	var intID int64
	params := mux.Vars(r)
	ID = params["bookId"]
	intID, _ = strconv.ParseInt(ID, 0, 0)
	// intID, _ = strconv.ParseInt(ID, 10, 64)

	bookDetails, _ := models.GetBookById(intID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// _ = json.NewDecoder(r.Body).Decode(&NewBook)
	createBook := &models.Book{}
	// models.Init()
	utils.ParseBody(r, createBook)

	book := createBook.CreateBook()
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var ID string
	var intID int64
	params := mux.Vars(r)
	ID = params["bookId"]
	intID, _ = strconv.ParseInt(ID, 0, 0)

	book := models.DeleteBook(intID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	var ID string
	var intID int64
	params := mux.Vars(r)
	ID = params["bookId"]
	intID, _ = strconv.ParseInt(ID, 0, 0)
	utils.ParseBody(r, updateBook)

	book, db := models.GetBookById(intID)

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}
	book.UpdatedAt = time.Now()
	db.Save(&book)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
