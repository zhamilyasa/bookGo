package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rest-project/internal/auth"
	"rest-project/internal/delivery"
	"rest-project/internal/middleware"

	"rest-project/internal/repository"
	service "rest-project/internal/services"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo, db)
	bookHandler := delivery.NewBookHandler(bookService)

	books := r.Group("api/v1/books")
	books.Use(middleware.AuthRequired())

	{
		books.GET("/", bookHandler.GetAllBooks)
		books.GET("/:id", bookHandler.GetBook)
		books.POST("/", bookHandler.CreateBook)
		books.PUT("/:id", bookHandler.UpdateBook)
		books.DELETE("/:id", bookHandler.DeleteBook)
		books.POST("/:id/assign", bookHandler.AddBookToUser)
	}

	authGroup := r.Group("/api/v1/auth")
	{
		authGroup.POST("/register", auth.Register)
		authGroup.POST("/login", auth.Login)
	}
}
