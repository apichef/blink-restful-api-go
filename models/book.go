package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(100)"`
	Description string `json:"description" gorm:"type:varchar(255)"`
	Authors []Author `json:"authors" gorm:"many2many:author_book;"`
	Genre Genre `json:"genre" gorm:"foreignKey:GenreID"`
	GenreID uint
	Publisher Publisher `json:"publisher" gorm:"foreignKey:PublisherID"`
	PublisherID uint
}