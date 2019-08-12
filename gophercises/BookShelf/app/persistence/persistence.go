package persistence

import (
	"BookShelf/app/helpers/logger"
	"BookShelf/app/models"
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbInstance *gorm.DB

//InitPostgresConn - Connection error
func InitPostgresConn() (*gorm.DB, error) {

	db, connectionError := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=BookShelf password=password")
	if connectionError != nil {
		logger.LogError("DB Connection error : ", connectionError)
		return nil, connectionError
	}

	//Register new models here for auto creation
	db.AutoMigrate(&models.Book{})

	dbInstance = db

	return dbInstance, nil
}

//GetDBInstance - Get database connection
func GetDBInstance() (*gorm.DB, error) {

	if dbInstance == nil {
		logger.LogError("DB not initialized")
		return nil, errors.New("DB not initialized")
	}
	return dbInstance, nil
}

//SetDBInstance - Used for test cases to initialize persistent layer
func SetDBInstance(db *gorm.DB) {
	dbInstance = db
}
