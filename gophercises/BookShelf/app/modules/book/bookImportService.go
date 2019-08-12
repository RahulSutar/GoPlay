package book

import (
	"BookShelf/app/helpers/logger"
	"BookShelf/app/models"
	"BookShelf/app/persistence"
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var bookListChannel = make(chan []models.Book, 10)
var mutex = &sync.Mutex{}
var wg sync.WaitGroup

//BulkImportFromCSVService - Import books from a CSV file
func BulkImportFromCSVService() (bool, error) {

	fileList, fileListError := GetCSVFileList()
	if fileListError != nil {
		logger.LogError("ImportFromCSVService")
		return false, nil
	}

	wg.Add(len(fileList))

	go func() {
		for _, importFilePath := range fileList {
			go ReadBooksFromCSVFile(bookListChannel, importFilePath)
		}
	}()

	go func() {
		for range fileList {
			go ImportBookList(bookListChannel, &wg)
		}
	}()

	wg.Wait()

	return true, nil
}

//GetCSVFileList - Get list of all csv files
func GetCSVFileList() ([]string, error) {

	var fileList []string

	dirScanError := filepath.Walk("./importDir/", func(path string, info os.FileInfo, err error) error {
		lowercasePath := strings.ToLower(path)
		if strings.HasSuffix(lowercasePath, ".csv") {
			fileList = append(fileList, path)
		}
		return nil
	})

	if dirScanError != nil {
		logger.LogError("GetCSVFileList dirScanError : ", dirScanError)
		return fileList, dirScanError
	}

	return fileList, nil
}

//ReadBooksFromCSVFile - Reads the CSV file and returns bookList
func ReadBooksFromCSVFile(bookListChannel chan []models.Book, filePath string) (int, []models.Book, error) {

	bookList := []models.Book{}
	recordCount := 0

	csvFile, fileError := os.Open(filePath)
	if fileError != nil {
		logger.LogError("ReadBooksCSVFile fileError : ", fileError)
		return recordCount, bookList, fileError
	}
	defer csvFile.Close()

	fileRecords, fileRecordsError := csv.NewReader(csvFile).ReadAll()
	if fileRecordsError != nil {
		logger.LogError("ReadBooksCSVFile fileRecordsError : ", fileRecordsError)
		return recordCount, bookList, fileRecordsError
	}

	for _, fileRecord := range fileRecords {
		bookData := models.Book{
			BookName:        strings.TrimSpace(fileRecord[0]),
			Author:          strings.TrimSpace(fileRecord[1]),
			BookType:        strings.TrimSpace(fileRecord[2]),
			BookDescription: strings.TrimSpace(fileRecord[3]),
		}
		bookList = append(bookList, bookData)
		recordCount++
	}

	if recordCount > 0 {
		processedBookList := PorcessBookList(bookList)
		bookListChannel <- processedBookList
	}

	logger.LogInfo(filePath, " read completed")

	return recordCount, bookList, nil

}

//PorcessBookList - Reads the booklist and aggregate books data by number of copies
func PorcessBookList(bookList []models.Book) []models.Book {

	processedBooksData := []models.Book{}

	for _, book := range bookList {

		bookFound := false

		// for _, processedBook := range processedBooksData {
		for i := 0; i < len(processedBooksData); i++ {
			if processedBooksData[i].BookName == book.BookName {
				bookFound = true
				processedBooksData[i].NumberOfCopies++
				break
			}
		}

		if !bookFound {
			book.NumberOfCopies = 1
			processedBooksData = append(processedBooksData, book)
		}

	}

	return processedBooksData
}

//ImportBookList - Imports a book list in DB
func ImportBookList(bookListChannel chan []models.Book, wg *sync.WaitGroup) (bool, error) {

	bookList := <-bookListChannel

	db, dbInstanceError := persistence.GetDBInstance()
	if dbInstanceError != nil {
		logger.LogError("ImportBookList dbInstanceError : ", dbInstanceError)
		return false, dbInstanceError
	}

	mutex.Lock()
	for _, bookToImport := range bookList {

		existingBook, existingBookError := GetBookByBookNameService(bookToImport.BookName)
		// if existingBookError != nil {
		// 	logger.LogError("ImportBookList existingBookError : ", existingBookError)
		// }

		if existingBookError != nil && existingBook.BookName == "" {
			insertBook := db.Create(&bookToImport)
			if insertBook.Error != nil {
				logger.LogError("ImportBookList insertError : '", bookToImport.BookName, "'", insertBook.Error)
			}
		} else {
			_, updateError := UpdateCountToExistingBookService(bookToImport.BookName, bookToImport.NumberOfCopies)
			if updateError != nil {
				logger.LogError("ImportBookList updateError : '", bookToImport.BookName, "'", updateError.Error)
			}
		}
	}
	logger.LogInfo("Write completed")
	mutex.Unlock()
	wg.Done()

	return true, nil
}
