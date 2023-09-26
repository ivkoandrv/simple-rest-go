package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	db *gorm.DB
}

func ConnectToDb() (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s  dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
	)
	fmt.Println("DSN: ", dsn)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{db: database}, nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}

func (d *Database) GetDBName() string {
	return os.Getenv("DB_NAME")
}
