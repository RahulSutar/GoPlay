package book

import (
	"BookShelf/app/persistence"
	"testing"
)

func setupTestCase(t *testing.T) func(t *testing.T) {
	db, dbError := persistence.InitPostgresConn()
	if dbError != nil {
		t.Error("Error setting up DB : ", dbError)
	}

	persistence.SetDBInstance(db)

	return func(t *testing.T) {
		db.Close()
	}
}

func TestGetAllBooksService(t *testing.T) {

	testSetup := setupTestCase(t)
	defer testSetup(t)

	_, bookListError := GetAllBooksService()

	if bookListError != nil {
		t.Error("Testing Error : ", bookListError)
		t.Fail()
	}

}

func TestUpdateCountToExistingBookService(t *testing.T) {

	testSetup := setupTestCase(t)
	defer testSetup(t)

	bookName := "HowToCodeInGO"

	existingBook, existingBookError := GetBookByBookNameService(bookName)
	if existingBookError != nil {
		t.Error("Testing Error : ", existingBookError)
		t.Fail()
	}

	updatedBook, updateError := UpdateCountToExistingBookService(bookName, 1)
	if updateError != nil {
		t.Error("Testing Error : ", updateError)
		t.Fail()
	}

	if updatedBook.NumberOfCopies == existingBook.NumberOfCopies {
		t.Fail()
	}

}

func TestDeleteBookInfoService(t *testing.T) {

	testSetup := setupTestCase(t)
	defer testSetup(t)

	bookName := "HowToCodeInGO"

	existingBook, existingBookError := GetBookByBookNameService(bookName)
	if existingBookError != nil {
		t.Error("Testing Error : ", existingBookError)
		t.Fail()
	}

	deletedBook, updateError := DeleteBookInfoService(bookName)
	if updateError != nil {
		t.Error("Testing Error : ", updateError)
		t.Fail()
	}

	if existingBook.NumberOfCopies == deletedBook.NumberOfCopies {
		t.Fail()
	}

}

func TestGetBookByBookNameService(t *testing.T) {

	testSetup := setupTestCase(t)
	defer testSetup(t)

	bookName := "HowToCodeInGO"

	existingBook, existingBookError := GetBookByBookNameService(bookName)
	if existingBookError != nil {
		t.Error("Testing Error : ", existingBookError)
		t.Fail()
	}

	if existingBook.BookName != bookName {
		t.Fail()
	}

}

func TestGetMockData(t *testing.T) {

	bookName := "HowToCodeInGO"

	existingBook, existingBookError := GetMockData(bookName)
	if existingBookError != nil {
		t.Error("Testing Error : ", existingBookError)
		t.Fail()
	}

	if existingBook.BookName != bookName {
		t.Fail()
	}

}
