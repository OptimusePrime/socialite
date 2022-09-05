package services

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"socialite/models"
)

func CreateUser(db *gorm.DB, user models.User) error {
	return db.Create(&user).Error
}

func FindUserByUUID(db *gorm.DB, uuid uuid.UUID) (models.User, error) {
	var user models.User
	err := db.First(&user, uuid).Error
	return user, err
}

func FindAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	err := db.Model(&models.User{}).Find(&users).Error
	return users, err
}

func DeleteOneUser(db *gorm.DB, uuid uuid.UUID) error {
	return db.Delete(&models.User{}, uuid).Error
}

func UpdateOneUser(db *gorm.DB, update models.User, uuid uuid.UUID) error {
	return db.Model(&models.User{}).Where("id = ?", uuid).Updates(update).Error
}
