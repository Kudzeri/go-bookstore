package models

import (
	"fmt"
	"github.com/kudzeri/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
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
