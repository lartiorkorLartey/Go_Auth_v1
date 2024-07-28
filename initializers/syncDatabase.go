package initializers

import "authapp.com/m/models"


func SyncDatabase() {
	DB.AutoMigrate(&models.Client{}, &models.User{})
}