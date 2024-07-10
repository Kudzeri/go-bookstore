package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=localhost user=postgres password=root dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Almaty"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	db = d
	fmt.Println("Database connection successful!")
}

func GetDB() *gorm.DB {
	return db
}
