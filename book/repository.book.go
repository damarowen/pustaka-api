package book

import (
	"pustaka-api/models"

	"gorm.io/gorm"
)

type Irepository interface {
	FindAll() ([]models.Book, error)
	FindById(ID uint) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(book models.Book) (models.Book, error)
}

type PustakaApiRepository struct {
	pustaka_api *gorm.DB
}

func NewBookRepository(db *gorm.DB) Irepository {
	return &PustakaApiRepository{pustaka_api: db}
}

func (r *PustakaApiRepository) FindAll() ([]models.Book, error) {
	var books []models.Book

	err := r.pustaka_api.Preload("User").Find(&books).Error

	return books, err
}

func (r *PustakaApiRepository) FindById(ID uint) (models.Book, error) {
	var book models.Book

	err := r.pustaka_api.Preload("User").Find(&book, ID).Error

	return book, err
}

func (r *PustakaApiRepository) Create(book models.Book) (models.Book, error) {
	err := r.pustaka_api.Save(&book).Error

	//* tampilin user dengan relasi ke book
	r.pustaka_api.Preload("User").Find(&book)
	return book, err
}

func (r *PustakaApiRepository) Update(book models.Book) (models.Book, error) {

	err := r.pustaka_api.Save(&book).Error
	r.pustaka_api.Preload("User").Find(&book)

	return book, err
}

func (r *PustakaApiRepository) Delete(book models.Book) (models.Book, error) {

	err := r.pustaka_api.Delete(&book).Error

	return book, err
}
