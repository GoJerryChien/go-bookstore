package models

import (
	"github.com/jinzhu/gorm"
	"github.com/GoJerryChien/go-bookstore/package/config"
)

var db *gorm.DB

type book struct{
	gorm.model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init()[
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
]

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []book
	db.Find(&Books)
	return Books
	
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook (ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}