package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string `gorm:"column:book_name"`
	Author      string
	Publication string
	Year        int
}
