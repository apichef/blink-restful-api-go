package dtos

type CreateBookDto struct {
	Name string `json:"name" binding:"required,max=100"`
	Description string `json:"description" binding:"max=300"`
	GenreID uint `json:"genre_id" binding:"required"`
	PublisherID uint `json:"publisher_id" binding:"required"`
}
