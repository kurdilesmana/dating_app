package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kurdilesmana/dating_app/config"
	"github.com/kurdilesmana/dating_app/internal/domain"
)

var DB *gorm.DB

func InitDatabase() {
	// Configure and open database connection
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword))

	if err != nil {
		panic("Failed to connect to the database")
	}

	// Migrate database tables
	db.AutoMigrate(&domain.User{})

	// Assign the database instance to the global variable
	DB = db
}
