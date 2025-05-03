package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func InitDB() {
	// Подгружаем .env
	_ = godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslmode := "disable"

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, sslmode)

	sqlDB, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("SQL connection failed:", err)
	}

	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatal("Migration driver init failed:", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/db/migrations", "postgres", driver)
	if err != nil {
		log.Fatal("Migration init failed:", err)
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal("Migration failed:", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("GORM connection failed:", err)
	}

	DB = gormDB
	log.Println("GORM connection successful")
}
