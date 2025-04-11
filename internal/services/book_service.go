package service

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
	"time"
)

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetById(id int) (*models.Book, error)
	Create(book *models.Book) error
	Update(id int, book *models.BookEdit) error
	Delete(bookID int) error
	GetFilteredBooks(author, sort, search string) ([]models.Book, error)
}

type BookService struct {
	repo BookRepository
	db   *gorm.DB
}

func NewBookService(bookRepo BookRepository, db *gorm.DB) *BookService {
	return &BookService{
		repo: bookRepo,
		db:   db,
	}
}

func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetBooksFiltered(author, sort, search string) ([]models.Book, error) {
	return s.repo.GetFilteredBooks(author, sort, search)
}

func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	return s.repo.GetById(id)
}

func (s *BookService) Create(title, author, publishedAt string, pages int) (*models.Book, error) {
	book := &models.Book{
		Title:       title,
		Author:      author,
		Pages:       pages,
		PublishedAt: time.Now(),
	}
	err := s.repo.Create(book)
	return book, err
}
func (s *BookService) AddBookToUser(userID uint, bookID uint) error {
	var user models.User
	if err := s.db.Preload("Books").First(&user, userID).Error; err != nil {
		return err
	}

	var book models.Book
	if err := s.db.First(&book, bookID).Error; err != nil {
		return err
	}

	return s.db.Model(&user).Association("Books").Append(&book)
}

func (s *BookService) Update(id int, bookEdit *models.BookEdit) (*models.Book, error) {
	err := s.repo.Update(id, bookEdit)
	if err != nil {
		return nil, err
	}
	return s.GetBookByID(id)
}

func (s *BookService) DeleteBook(bookID int) error {
	return s.repo.Delete(bookID)
}
