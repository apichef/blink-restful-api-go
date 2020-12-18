package models

import (
	"gorm.io/gorm"
)

type Publisher struct {
	gorm.Model
	Name string `json:"name"`
	Books []Book `json:"books" gorm:"foreignKey:PublisherID"`
}