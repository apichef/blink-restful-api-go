package models

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(100)"`
	Books []Book `json:"books" gorm:"foreignKey:GenreID"`
}