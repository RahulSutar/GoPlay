package test

import "BookShelf/app/models"

func GetBookData(bookService BookServiceInterface, bookName string) (models.Book, error) {
	return bookService.getBookData(bookName)
}
