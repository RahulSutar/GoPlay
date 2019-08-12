package routes

import (
	"BookShelf/app/modules/book"

	"github.com/gin-gonic/gin"
)

//InitBookRoutes - APIs related to books
func InitBookRoutes(bookAPI *gin.RouterGroup) {

	bookAPI.GET("/getallbooks", book.GetAllBooksAPI)
	bookAPI.GET("/getbook/:bookName", book.GetBookByBookNameAPI)
	bookAPI.POST("/newbook", book.AddNewBookAPI)
	bookAPI.POST("/updatebookinfo", book.UpdateBookInfoAPI)
	bookAPI.POST("/deletebook", book.DeleteBookInfoAPI)
	bookAPI.POST("/importfromcsv", book.ImportFromCSVAPI)
	// bookAPI.POST("/testapi", book.TestAPI)

}
