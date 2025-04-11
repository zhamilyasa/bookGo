package db

import (
	"database/sql"
	"fmt"
	"log"
	"rest-project/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func InitDB() {
	dbHost := "localhost"
	dbName := "mydatabase"
	dbUser := "myuser"
	dbPass := "mypassword"
	dbPort := "5432" // Проверь что порт совпадает с твоим postgres
	sslmode := "disable"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)

	sqlDB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/db/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&models.Book{}, &models.User{})
	if err != nil {
		log.Fatal("Error migrating:", err)
	} else {
		log.Println("Migration successful")
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = gormDB
}
