package models

import (
	"fmt"
	"github.com/kudzeri/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()

	err := db.AutoMigrate(&Book{})
	if err != nil {
		fmt.Printf("failed to migrate:%v", err)
	}
}
func (b *Book) CreateBook() *Book {
	err := db.Create(&b).Error
	if err != nil {
		fmt.Printf("failed to create book: %v", err)
		return nil
	}

	return b
}

func GetAllBook() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)

	return &getBook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}

func (b *Book) UpdateBook() *Book {
	err := db.Save(b).Error
	if err != nil {
		fmt.Printf("models.UpdateBook(): %v\n", err)
		return nil
	}
	return b
}
