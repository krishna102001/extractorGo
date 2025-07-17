package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializedDB() {
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_ssl_mode := os.Getenv("DB_SSL_MODE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s channel_binding=require", db_host, db_user, db_password, db_name, db_port, db_ssl_mode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the postgres database %s", err.Error())
	}
	DB = db
	log.Println("Successfully Connected to Database!!")
	migrateDB()
}

func migrateDB() {
	if err := DB.AutoMigrate(&ExtractsTable{}, &ConvertsTable{}); err != nil {
		log.Println("Failed to migrate the Database ", err.Error())
	}
	log.Println("Successfully migrated the table")
}
