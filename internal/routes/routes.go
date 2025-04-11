package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rest-project/internal/delivery"
	"rest-project/internal/repository"

	service "rest-project/internal/services"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := delivery.NewBookHandler(bookService)

	books := r.Group("api/v1/books")
	{
		books.GET("/", bookHandler.GetAllBooks) // Уже включает фильтрацию через query-параметры
		// books.GET("/books/filter", bookHandler.GetFilteredBooks) ← удалить или закомментировать
		books.GET("/:id", bookHandler.GetBook)
		books.POST("/", bookHandler.CreateBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
	}

}
