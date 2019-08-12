package book

import (
	"BookShelf/app/helpers/logger"
	"BookShelf/app/models"
	"BookShelf/app/persistence"
	"errors"
	"strings"
)

//GetAllBooksService - Service to get all books
func GetAllBooksService() ([]models.Book, error) {

	books := []models.Book{}

	db, dbInstanceError := persistence.GetDBInstance()
	if dbInstanceError != nil {
		logger.LogError("AddNewBook dbInstanceError : ", dbInstanceError)
		return books, dbInstanceError
	}

	findBooksResult := db.Find(&books)
	if findBooksResult.Error != nil {
		logger.LogError("GetAllBooksService resultError : ", findBooksResult.Error)
		return books, findBooksResult.Error
	}

	return books, nil
}

//AddNewBookService - Service to add a new book to DB
func AddNewBookService(newBook models.Book) (models.Book, error) {

	db, dbInstanceError := persistence.GetDBInstance()
	if dbInstanceError != nil {
		logger.LogError("AddNewBook dbInstanceError : ", dbInstanceError)
		return newBook, dbInstanceError
	}

	existingBook, findExistingBookError := GetBookByBookNameService(newBook.BookName)
	if findExistingBookError != nil {
		logger.LogError("AddNewBook findExistingBookError : ", findExistingBookError)
	}

	if existingBook.BookName == "" {

		insertBook := db.Create(&newBook)
		if insertBook.Error != nil {
			logger.LogError("AddNewBook insertError : ", insertBook.Error)
			return newBook, insertBook.Error
		}

		return newBook, nil

	}

	existingBook, addCountError := UpdateCountToExistingBookService(existingBook.BookName, 1)
	if addCountError != nil {
		logger.LogError("AddNewBook addCountError : ", addCountError)
		return newBook, addCountError
	}

	return existingBook, nil

}

//UpdateCountToExistingBookService - Increment count of the given book
func UpdateCountToExistingBookService(bookName string, countToAdd int64) (models.Book, error) {

	existingBook := models.Book{}

	existingBook, findBookError := GetBookByBookNameService(bookName)
	if findBookError != nil {
		logger.LogError("UpdateCountToExistingBookService findBookError : ", findBookError)
		return existingBook, findBookError
	}

	if (existingBook.NumberOfCopies + countToAdd) >= 0 {
		existingBook.NumberOfCopies = existingBook.NumberOfCopies + countToAdd
	} else {
		existingBook.NumberOfCopies = 0
	}

	db, dbInstanceError := persistence.GetDBInstance()
	if dbInstanceError != nil {
		logger.LogError("UpdateCountToExistingBookService dbInstanceError : ", dbInstanceError)
		return existingBook, dbInstanceError
	}

	countUpdateError := db.Save(&existingBook)

	if countUpdateError.Error != nil {
		logger.LogError("UpdateCountToExistingBookService countUpdateError : ", countUpdateError.Error)
		return existingBook, countUpdateError.Error
	}

	return existingBook, nil
}

//GetBookByBookNameService - Get a book from db when name is provided
func GetBookByBookNameService(bookName string) (models.Book, error) {

	bookDetails := models.Book{}

	bookName = strings.TrimSpace(bookName)
	if bookName == "" {
		logger.LogError("GetBookByBookName : bookName required")
		return bookDetails, errors.New("GetBookByBookName : bookName required")
	}

	db, dbInstanceError := persistence.GetDBInstance()
	if dbInstanceError != nil {
		logger.LogError("GetBookByBookNameService dbInstanceError : ", dbInstanceError)
		return bookDetails, dbInstanceError
	}

	result := db.Where("book_name = ?", bookName).First(&bookDetails)
	if result.Error != nil {
		logger.LogError("GetBookByBookNameService dbError: ", result.Error)
		return bookDetails, errors.New("GetBookByBookNameService dbError: " + result.Error.Error())
	}
	return bookDetails, nil
}

//UpdateBookInfoService - Update book information (author, bookType and description only)
func UpdateBookInfoService(bookInfo models.Book) (models.Book, error) {

	existingBook, existingBookError := GetBookByBookNameService(bookInfo.BookName)
	if existingBookError != nil {
		logger.LogError("UpdateBookInfoService existingBookError : ", existingBookError)
		return existingBook, existingBookError
	}

	existingBook.Author = bookInfo.Author
	existingBook.BookType = bookInfo.BookType
	existingBook.BookDescription = bookInfo.BookDescription

	db, dbInstanceError := persistence.GetDBInstance()
	if dbInstanceError != nil {
		logger.LogError("UpdateBookInfoService dbInstanceError : ", dbInstanceError)
		return existingBook, dbInstanceError
	}

	infoUpdateError := db.Save(&existingBook)

	if infoUpdateError.Error != nil {
		logger.LogError("UpdateBookInfoService infoUpdateError : ", infoUpdateError.Error)
		return existingBook, infoUpdateError.Error
	}

	return existingBook, nil
}

//DeleteBookInfoService - Delete book (reduce book count)
func DeleteBookInfoService(bookName string) (models.Book, error) {
	return UpdateCountToExistingBookService(bookName, -1)
}

//GetMockData - Mocks a DB call
func GetMockData(bookName string) (models.Book, error) {

	bookName = strings.TrimSpace(bookName)

	bookObject := models.Book{}

	if bookName == "" {
		logger.LogError("GetMockData : bookName required")
		return bookObject, errors.New("GetMockData : bookName required")
	}

	if bookName != "HowToCodeInGO" {
		logger.LogError("GetMockData : bookNotFound")
	}

	bookObject.BookName = "HowToCodeInGO"
	bookObject.Author = "Test Author"
	bookObject.BookType = "Technical book for a test"
	bookObject.BookDescription = "This is the test description for the book"

	return bookObject, nil

}
