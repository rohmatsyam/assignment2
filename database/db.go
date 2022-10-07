package database

import (
	"assignment2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "rohmat_syam"
	dbname   = "assignment-dua"
	db   *gorm.DB
	err      error
)

func GetConnection() *gorm.DB {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db,err = gorm.Open(postgres.Open(config),&gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database : ",err)		
	}	
	log.Println("DB Connecting Established...")
	db.AutoMigrate(&models.Item{})	
	db.AutoMigrate(&models.Order{})	
	return db
}