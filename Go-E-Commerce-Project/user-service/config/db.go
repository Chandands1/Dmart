package config

import (
	"fmt"
	"user-service/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	
)

var DB *gorm.DB

func ConnectDB(){
	dsn := "host=localhost user=postgres password=chandan123 dbname=dmart port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		panic("failed to connect to database")
	}
	fmt.Println("Connected to the database")
	DB = db

    Migrate()
}

func Migrate() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf(" migration failed: %v", err)
	}
	fmt.Println(" Database migrated")
}
