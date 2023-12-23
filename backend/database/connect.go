package database

import (
	"fmt"
	"strconv"
	"log"

	"sm-system/configs"
	"sm-system/internals/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := configs.Config("POSTGRES_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Idiot")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configs.Config("POSTGRES_HOST"),
		port,
		configs.Config("POSTGRES_USER"),
		configs.Config("POSTGRES_PASSWORD"),
		configs.Config("POSTGRES_DB"))

	DB, err = gorm.Open(postgres.Open(dsn)) 
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connect to database")

	DB.AutoMigrate(&models.Sert{})
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database migrated")
}

