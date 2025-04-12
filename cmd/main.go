package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"rest-project/internal/auth"
	"rest-project/internal/db"
	"rest-project/internal/routes"
)

func main() {
	db.InitDB()
	database := db.DB
	auth.Init(database)

	r := gin.Default()
	routes.SetupRoutes(r, database)

	if err := r.Run(":8081"); err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}
}
