package service

import (
	"books/dto"
	"books/entity"
	"books/repository"
	"time"
)

type BookService interface {
	All() []entity.Book
	FindByID(id uint64) entity.Book
	Create(book dto.BookRequest) bool
	Update(id uint64, book dto.UpdateBookRequest) bool
	Delete(id uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (service *bookService) All() []entity.Book {
	return service.bookRepository.AllBook()
}

func (service *bookService) FindByID(id uint64) entity.Book {
	return service.bookRepository.FindByID(id)
}

func (service *bookService) Create(book dto.BookRequest) bool {
	var (
		newBook entity.Book
	)

	newBook.Title = book.Title
	newBook.Author = book.Author
	newBook.ISBN = book.ISBN
	publishedDate, _ := time.Parse("2006-01-02", book.PublishedDate)
	newBook.PublishedDate = publishedDate
	result := service.bookRepository.Create(newBook)

	return result
}

func (service *bookService) Update(id uint64, book dto.UpdateBookRequest) bool {
	var (
		newBook entity.Book
	)

	newBook.Title = book.Title
	newBook.Author = book.Author
	newBook.ISBN = book.ISBN
	publishedDate, _ := time.Parse("2006-01-02", book.PublishedDate)
	newBook.PublishedDate = publishedDate
	result := service.bookRepository.Update(id, newBook)

	if result != true {
		return false
	}
	return result
}

func (service *bookService) Delete(id uint64) bool {
	result := service.bookRepository.Delete(id)

	if result != true {
		return false
	}
	return result
}
