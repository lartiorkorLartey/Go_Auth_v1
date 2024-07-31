package initializers

import "github.com/InnocentEdem/goauth/models"



func SyncDatabase() {
	DB.AutoMigrate(&models.Client{}, &models.User{}, &models.FeatureRequest{})
}