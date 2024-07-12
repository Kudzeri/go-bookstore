package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kudzeri/go-bookstore/pkg/models"
	"github.com/kudzeri/go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var newBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, err := json.Marshal(b)
	if err != nil {
		fmt.Printf("bookController.CreateBook(): %v\n", err)
		http.Error(w, "Error marshalling book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBook()
	res, err := json.Marshal(newBooks)
	if err != nil {
		fmt.Printf("bookController.GetBooks(): %v\n", err)
		http.Error(w, "Error marshalling books", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Printf("bookController.GetBookByID(): %v\n", err)
	}
	bookDetails, _ := models.GetBookByID(ID)
	res, err := json.Marshal(bookDetails)
	if err != nil {
		fmt.Printf("bookController.GetBookByID(): %v\n", err)
		http.Error(w, "Error marshalling book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Printf("bookController.UpdateBook(): %v\n", err)
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	bookDetails, _ := models.GetBookByID(ID)

	bookDetails.Name = updateBook.Name
	bookDetails.Author = updateBook.Author
	bookDetails.Publication = updateBook.Publication

	b := bookDetails.UpdateBook()
	res, err := json.Marshal(b)
	if err != nil {
		fmt.Printf("bookController.UpdateBook(): %v\n", err)
		http.Error(w, "Error marshalling book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Printf("bookController.DeleteBook(): %v\n", err)
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book := models.DeleteBook(ID)
	res, err := json.Marshal(book)
	if err != nil {
		fmt.Printf("bookController.DeleteBook(): %v\n", err)
		http.Error(w, "Error marshalling book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
