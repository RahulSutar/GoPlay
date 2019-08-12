package book

import (
	"BookShelf/app/helpers/logger"
	"BookShelf/app/models"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//GetAllBooksAPI - API to get all the books
func GetAllBooksAPI(c *gin.Context) {
	bookList, serviceError := GetAllBooksService()
	if serviceError != nil {
		logger.LogError("GetAllBooksAPI serviceError : ", serviceError)
		c.JSON(http.StatusInternalServerError, serviceError)
	}
	c.JSON(http.StatusOK, bookList)
}

//GetBookByBookNameAPI - API to get all the books
func GetBookByBookNameAPI(c *gin.Context) {

	bookName := c.Param("bookName")
	bookName = strings.TrimSpace(bookName)

	if bookName == "" {
		logger.LogError("GetBookByBookNameAPI : bookName required")
		c.JSON(http.StatusInternalServerError, errors.New("GetBookByBookNameAPI : bookName required"))
		return
	}

	bookDetails, getBookServiceError := GetBookByBookNameService(bookName)
	if getBookServiceError != nil {
		logger.LogError("GetBookByBookNameAPI serviceError : ", getBookServiceError)
		c.JSON(http.StatusInternalServerError, getBookServiceError)
	}

	c.JSON(http.StatusOK, bookDetails)
	return
}

//AddNewBookAPI - API to get all the books
func AddNewBookAPI(c *gin.Context) {

	bookDetails := models.Book{}

	bindError := c.Bind(&bookDetails)
	if bindError != nil {
		logger.LogError("AddNewBookAPI bindError : ", bindError)
		c.JSON(http.StatusInternalServerError, bindError)
		return
	}

	newBookData, newBookServiceError := AddNewBookService(bookDetails)
	if newBookServiceError != nil {
		logger.LogError("AddNewBookAPI serviceError : ", newBookServiceError)
		c.JSON(http.StatusInternalServerError, newBookServiceError)
		return
	}

	c.JSON(http.StatusOK, newBookData)
	return
}

//UpdateBookInfoAPI - API to update book information
func UpdateBookInfoAPI(c *gin.Context) {

	bookInfo := models.Book{}

	bindError := c.Bind(&bookInfo)
	if bindError != nil {
		logger.LogError("UpdateBookInfoAPI bindError : ", bindError)
		c.JSON(http.StatusInternalServerError, bindError)
		return
	}

	updatedBook, bookInfoUpdateError := UpdateBookInfoService(bookInfo)
	if bookInfoUpdateError != nil {
		logger.LogError("UpdateBookInfoAPI bookInfoUpdateError : ", bookInfoUpdateError)
		c.JSON(http.StatusInternalServerError, bookInfoUpdateError)
		return
	}

	c.JSON(http.StatusOK, updatedBook)

	return
}

//DeleteBookInfoAPI - API to delete book (reduce available copies)
func DeleteBookInfoAPI(c *gin.Context) {

	bookInfo := models.Book{}

	bindError := c.Bind(&bookInfo)
	if bindError != nil {
		logger.LogError("DeleteBookInfoAPI bindError : ", bindError)
		c.JSON(http.StatusInternalServerError, bindError)
		return
	}

	deletedBook, deleteBookError := DeleteBookInfoService(bookInfo.BookName)
	if deleteBookError != nil {
		logger.LogError("DeleteBookInfoAPI deleteBookError : ", deleteBookError)
		c.JSON(http.StatusInternalServerError, deleteBookError)
		return
	}

	c.JSON(http.StatusOK, deletedBook)
	return
}

//ImportFromCSVAPI - API to import books from CSV files
func ImportFromCSVAPI(c *gin.Context) {

	importStatus, importError := BulkImportFromCSVService()
	if importError != nil {
		logger.LogError("ImportFromCSVAPI importError : ", importError)
		c.JSON(http.StatusOK, importError)
		return
	}

	c.JSON(http.StatusOK, importStatus)
	return
}

//TestAPI - API for testing functions
// func TestAPI(c *gin.Context) {
// 	recordCount, bookList, readError := ReadBooksFromCSVFile("./importDir/file1.csv")
// 	if readError != nil {
// 		logger.LogError("Test API error : ", readError)
// 	}

// 	fmt.Println("records count : ", recordCount)

// 	processedBooksData := PorcessBookList(bookList)

// 	c.JSON(http.StatusOK, processedBooksData)

// }
