package models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(100)"`
	Books []Book `json:"books" gorm:"many2many:author_book;"`
}
