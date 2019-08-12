package routes

import (
	"github.com/gin-gonic/gin"
)

// Init -Init
func Init(g *gin.Engine) {

	bookAPI := g.Group("/book")
	InitBookRoutes(bookAPI)

}
