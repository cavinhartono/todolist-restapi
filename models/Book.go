package models

import (
	"fmt"

	"github.com/cavinhartono/todolist-restapi/config"
	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID     int     `json:"id"`
	Author string  `json:"author"`
	Title  string  `json:"title"`
	Price  float64 `json:"price"`
	Isbn   string  `json:"isbn"`
}

func (b *Book) TableName() string {
	return "book"
}

func GetAllBooks(book *[]Book) (err error) {
	if err = config.DB.Find(book).Error; err != nil {
		return err
	}
	return nil
}

func CreateBook(book *Book) (err error) {
	if err = config.DB.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func GetBookByID(book *Book, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(book).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBook(book *Book, id string) (err error) {
	fmt.Println(book)
	config.DB.Save(book)
	return nil
}

func DeleteBook(book *Book, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(book)
	return nil
}
