package initializers

import "Authentication/models"

func SyncDataBase() {
	// Load environment variables
	DB.AutoMigrate(&models.User{})

}
