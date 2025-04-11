package repository

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}
type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetById(id int) (*models.Book, error)
	Create(book *models.Book) error
	Update(id int, book *models.BookEdit) error
	Delete(bookID int) error
	GetFilteredBooks(author, sort, search string) ([]models.Book, error) // ← вот это строка обязательна
}

func NewBookRepository(db *gorm.DB) *BookRepositoryImpl {
	return &BookRepositoryImpl{db: db}
}

func (r BookRepositoryImpl) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r BookRepositoryImpl) GetById(id int) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r BookRepositoryImpl) GetFilteredBooks(author, sort, search string) ([]models.Book, error) {
	var books []models.Book
	query := r.db

	if search != "" {
		query = query.Where("title ILIKE ? OR author ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if author != "" {
		query = query.Where("author = ?", author)
	}

	switch sort {
	case "title_asc":
		query = query.Order("title ASC")
	case "title_desc":
		query = query.Order("title DESC")
	case "date_asc":
		query = query.Order("published_at ASC")
	case "date_desc":
		query = query.Order("published_at DESC")
	}

	err := query.Find(&books).Error
	return books, err
}
func (r BookRepositoryImpl) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r BookRepositoryImpl) Update(id int, book *models.BookEdit) error {
	return r.db.Model(&models.Book{}).Where("id = ?", id).Omit("id").Updates(book).Error
}

func (r BookRepositoryImpl) Delete(bookID int) error {
	return r.db.Delete(&models.Book{}, bookID).Error
}
