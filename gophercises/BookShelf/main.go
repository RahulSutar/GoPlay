package main

import (
	"BookShelf/app/helpers/logger"
	"BookShelf/app/persistence"
	"BookShelf/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	g := gin.New()

	logger.InitLogging()

	routes.Init(g)

	_, persistentError := persistence.InitPostgresConn()
	if persistentError != nil {
		logger.LogError("DB persistentError : ", persistentError)
		return
	}

	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Run(":3032")

}
