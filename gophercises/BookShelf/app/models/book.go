package models

import (
	"github.com/jinzhu/gorm"
)

//Book - Properties of book
type Book struct {
	gorm.Model
	BookName           string `json:"bookName" gorm:"bookName"`
	Author             string `json:"author" gorm:"author"`
	BookType           string `json:"bookType" gorm:"bookType"`
	BookDescription    string `json:"bookDescription" gorm:"bookDescription"`
	NumberOfCopies     int64  `json:"numberOfCopies" gorm:"numberOfCopies"`
	AvailabilityStatus string `json:"availabilityStatus" gorm:"availabilityStatus"`
}

//BeforeCreate - GORM hook to process and set default number of copies before new insert
func (book *Book) BeforeCreate() (err error) {
	if book.NumberOfCopies == 0 {
		book.NumberOfCopies = 1
		book.AvailabilityStatus = "available"
	}
	return
}

//BeforeUpdate - GORM hook to process and set availability status based on number of copies
func (book *Book) BeforeUpdate() (err error) {
	if book.NumberOfCopies <= 0 {
		book.AvailabilityStatus = "unavailable"
	} else {
		book.AvailabilityStatus = "available"
	}
	return
}
