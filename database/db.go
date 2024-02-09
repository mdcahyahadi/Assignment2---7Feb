package database

import (
	"assignment-golang8-7feb/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "password"
	port     = "5432"
	dbname   = "assignment"
	db       *gorm.DB
	err      error
)

func GetDB() *gorm.DB {
	return db
}

func StartConnection() {
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}
