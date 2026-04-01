package db

import "go-api/internal/models"

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
}
