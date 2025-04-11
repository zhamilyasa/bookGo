package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"rest-project/internal/auth"
	"rest-project/internal/db"
	"rest-project/internal/models"
	"rest-project/internal/routes"
)

func main() {
	db.InitDB()

	database := db.DB

	if err := database.AutoMigrate(&models.Book{}, &models.User{}); err != nil {
		log.Fatal("Error on migrating models:", err)
	}

	auth.Init(database)

	r := gin.Default()
	routes.SetupRoutes(r, database)

	if err := r.Run(":8081"); err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}

	//db, err := gorm.Open(postgres.Open("postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable"), &gorm.Config{})
	////m, err := migrate.NewWithDatabaseInstance("file://internal/db/migrations", "postgres", driver)
	////sqlDB, err := sql.Open("postgres", dbUrl) // уже есть
	////driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	////m, err := migrate.NewWithDatabaseInstance("file://internal/db/migrations", "postgres", driver)
	//auth.Init(db)
	//
	//if err != nil {
	//	log.Fatal("Error connecting to the database:", err)
	//}
	//
	//err = db.AutoMigrate(&models.Book{})
	//if err != nil {
	//	log.Fatal("Error on migrating to the DB", err)
	//}
	//
	//r := gin.Default()
	//routes.SetupRoutes(r, db)
	//r.Run(":8081")
}
