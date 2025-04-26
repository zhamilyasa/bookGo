package service

import (
	"errors"
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

func (s *BookService) GetUserBooks(userID uint) ([]models.Book, error) {
	var user models.User
	if err := s.db.Preload("Books").First(&user, userID).Error; err != nil {
		return nil, err
	}
	return user.Books, nil
}

func (s *BookService) Create(userID uint, title, author, publishedAt string, pages int) (*models.Book, error) {
	book := &models.Book{
		Title:       title,
		Author:      author,
		Pages:       pages,
		PublishedAt: time.Now(),
		CreatorID:   userID, //
	}

	err := s.repo.Create(book)
	if err != nil {
		return nil, err
	}

	// Привязываем книгу к пользователю
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	if err := s.db.Model(&book).Association("Users").Append(&user); err != nil {
		return nil, err
	}

	return book, nil
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

func (s *BookService) Update(id int, bookEdit *models.BookEdit, userID uint) (*models.Book, error) {
	book, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	if book.CreatorID != userID {
		return nil, errors.New("you can only update your own books")
	}
	err = s.repo.Update(id, bookEdit)
	if err != nil {
		return nil, err
	}
	return s.GetBookByID(id)
}

func (s *BookService) DeleteBook(bookID int, userID uint) error {
	book, err := s.repo.GetById(bookID)
	if err != nil {
		return err
	}
	if book.CreatorID != userID {
		return errors.New("you can only delete your own books")
	}
	return s.repo.Delete(bookID)
}
