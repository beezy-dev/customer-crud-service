package inits

import "github.com/beezy-dev/traveler-crud-service/models"

func MigrateSchema() {
	LoadEnv()
	DB.AutoMigrate(&models.Traveler{})
}
