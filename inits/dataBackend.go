package inits

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {

	// loading variables from .env file
	LoadEnv()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// connecting to database
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPass, dbName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// handling connection error
	if err != nil {
		log.Println("Error: failed to connect to database.")
		log.Println("Creating database")

		stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", dbName)
		rs := DB.Raw(stmt)
		if rs.Error != nil {
			log.Println("Database ", dbName, " exists.")
		} else {
			stmt := fmt.Sprintf("CREATE DATABASE '%s';", dbName)
			rs := DB.Exec(stmt)
			if rs.Error != nil {
				log.Println("Error: failed to create database:", dbName)
			}
		}
	} else {
		log.Println("Database", dbName, "exists.")
	}
}
