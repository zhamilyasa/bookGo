package delivery

import (
	"github.com/gin-gonic/gin"
	_ "log"
	"net/http"
	"rest-project/internal/models"
	service "rest-project/internal/services"
	"strconv"
)

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

type BookHandler struct {
	service *service.BookService
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, _ := h.service.GetAllBooks()
	c.JSON(http.StatusOK, books)
}
func (h *BookHandler) GetFilteredBooks(c *gin.Context) {
	author := c.Query("author")
	sort := c.Query("sort")
	search := c.Query("search")

	books, err := h.service.GetBooksFiltered(author, sort, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to filter books"})
		return
	}

	c.JSON(http.StatusOK, books)
}
func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := h.service.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) AddBookToUser(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = h.service.AddBookToUser(userID, uint(bookID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book added to user"})
}

func (h *BookHandler) GetUserBooks(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	books, err := h.service.GetUserBooks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) CreateBook(c *gin.Context) {

	userID := c.MustGet("userID").(uint)

	var bookCreate models.BookEdit
	if err := c.ShouldBindJSON(&bookCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newBook, err := h.service.Create(userID, bookCreate.Title, bookCreate.Author, bookCreate.PublishedAt, bookCreate.Pages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, newBook)

}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	userID := c.MustGet("userID").(uint) // Получаем ID пользователя из контекста

	var bookEdit models.BookEdit
	if err := c.ShouldBindJSON(&bookEdit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedBook, err := h.service.Update(id, &bookEdit, userID) // Передаем userID в сервис
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()}) // Возвращаем ошибку, если нет прав
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	userID := c.MustGet("userID").(uint) // Получаем ID пользователя из контекста

	if err := h.service.DeleteBook(id, userID); err != nil { // Передаем userID в сервис
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()}) // Возвращаем ошибку, если нет прав
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
