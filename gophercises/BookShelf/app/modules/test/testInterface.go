package test

import (
	"BookShelf/app/models"
	"BookShelf/app/modules/book"
	"errors"
	"strings"
)

//BookServiceInterface - BooksInterface
type BookServiceInterface interface {
	getBookData(bookName string) (models.Book, error)
}

type ActualData struct{}
type MockData struct{}

func (ad ActualData) getBookData(bookName string) (models.Book, error) {

	bookName = strings.TrimSpace(bookName)
	bookObj := models.Book{}

	if bookName == "" {
		return bookObj, errors.New("bookName required !!")
	}

	bookData, bookDataError := book.GetBookByBookNameService(bookName)
	if bookDataError :=  nil {
		return bookObj, bookDataError
	}

	return bookObj, nil
}


func (md MockData) getBookData(bookName string) (models.Book, error) {

	bookName = strings.TrimSpace(bookName)
	bookObj := models.Book{}

	if bookName == "" {
		return bookObj, errors.New("bookName required !!")
	}

	bookObj.BookName = "TestBook"

	return bookObj, nil
}