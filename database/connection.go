package database

import (
	"fmt"
	"github.com/apichef/blink-restful-api-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Open() *gorm.DB {
	db, err := gorm.Open(postgres.Open(getConnectionString()), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}

func Migrate() {
	db := Open()

	err := db.AutoMigrate(
		&models.Author{},
		&models.Genre{},
		&models.Publisher{},
		&models.Book{},
	)

	if err != nil {
		panic(err.Error())
	}
}

func getConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Amsterdam",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "dbuser"),
		getEnv("DB_PASS", "dbpass"),
		getEnv("DB_NAME", "blink"),
		getEnv("DB_PORT", "5432"),
	)
}

func getEnv(key string, defaultValue string) string  {
	val := os.Getenv(key)

	if val != "" {
		return val
	}

	return defaultValue
}

func Close(db *gorm.DB) {
	sqlDb, err := db.DB()

	if err != nil {
		panic("Failed to connect to database")
	}

	sqlDb.Close()
}
