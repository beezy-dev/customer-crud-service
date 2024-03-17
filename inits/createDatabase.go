package inits

import (
	"log"
	"os"
)

func CreateDBPostgres() {

	var err error

	LoadEnv()

	dbType := os.Getenv("DB_TYPE")
	dbName := os.Getenv("DB_NAME")

	// check if db exists
	log.Println("Info: checking if database", dbName, "exists in service", dbType)
	rs := DB.Exec("SELECT datname FROM pg_catalog.pg_database WHERE datname='" + dbName + "';")

	if rs.RowsAffected != 1 {
		log.Println("Warning: database", dbName, "does not exist in service", dbType)
		log.Println("Info: creating database", dbName, "exists in service", dbType)
		rs := DB.Exec("CREATE DATABASE " + dbName + ";")
		if rs.Error != nil {
			log.Println("Error: can not create database ", dbName, "in service ", dbType)
			log.Fatal("Fatal: ", rs.Error)
		}
		log.Println("Info: checking if database", dbName, "exists now in service", dbType)
		rs = DB.Exec("SELECT datname FROM pg_catalog.pg_database WHERE datname='" + dbName + "';")
		if rs.RowsAffected != 1 {
			log.Println("Error: database ", dbName, " does not exist after create (query returned: ", rs.RowsAffected, "row.")
			log.Fatal("Fatal: ", rs.Error)
		} else {
			log.Println("Info: database", dbName, "exists in service", dbType)
		}
	} else {
		log.Println("Info: database", dbName, "exists in service", dbType)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Println(err)
	}
	sqlDB.Close()

}
