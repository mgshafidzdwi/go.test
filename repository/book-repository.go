package repository

import (
	"books/entity"

	"gorm.io/gorm"
)

type BookRepository interface {
	AllBook() []entity.Book
	FindByID(id uint64) entity.Book
	Create(book entity.Book) bool
	Update(id uint64, book entity.Book) bool
	Delete(id uint64) bool
}

type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConn,
	}
}

func (db *bookConnection) AllBook() []entity.Book {
	var books []entity.Book
	db.connection.Find(&books)
	return books
}

func (db *bookConnection) FindByID(id uint64) entity.Book {
	var book entity.Book
	db.connection.Find(&book, id)
	return book
}

func (db *bookConnection) Create(book entity.Book) bool {
	result := db.connection.Save(&book)
	if result.Error != nil {
		return false
	}
	return true
}

func (db *bookConnection) Update(id uint64, book entity.Book) bool {
	var oldBook entity.Book
	db.connection.Find(&oldBook, id)

	oldBook.Title = book.Title
	oldBook.Author = book.Author
	oldBook.ISBN = book.ISBN
	oldBook.PublishedDate = book.PublishedDate

	result := db.connection.Save(&oldBook)
	if result.Error != nil {
		return false
	}
	return true
}

func (db *bookConnection) Delete(id uint64) bool {
	var book entity.Book
	db.connection.Delete(&book, id)
	return true
}
