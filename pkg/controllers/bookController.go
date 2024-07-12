package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kudzeri/go-bookstore/pkg/models"
	"net/http"
)

var newBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {

}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBook()
	res, err := json.Marshal(newBooks)
	if err != nil {
		fmt.Printf("bookController.GetBooks(): %v\n", err)
		http.Error(w, "Error marshalling books", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
