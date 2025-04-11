package service

import (
	"rest-project/internal/models"
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
}

func NewBookService(bookRepo BookRepository) *BookService {
	return &BookService{repo: bookRepo}
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
		PublishedAt: publishedAt,
		Pages:       pages,
	}
	err := s.repo.Create(book)
	return book, err
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
