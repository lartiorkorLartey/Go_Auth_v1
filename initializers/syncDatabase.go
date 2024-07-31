package initializers

import "github.com/InnocentEdem/Go_Auth_v1/models"



func SyncDatabase() {
	DB.AutoMigrate(&models.Client{}, &models.User{}, &models.FeatureRequest{})
}