package service

import (
	"gogin/config"
	"gogin/models"
)

// Get all users
func GetAllUsers() ([]models.User, error) {
    var users []models.User
    err := config.DB.Find(&users).Error
    return users, err
}

// Get user by ID
func GetUserByID(id int) (models.User, error) {
    var user models.User
    err := config.DB.First(&user, id).Error
    return user, err
}

// Create a new user
func CreateUser(user *models.User) error {
    return config.DB.Create(user).Error
}

// Update a user
func UpdateUser(id int, user *models.User) error {
    return config.DB.Model(&user).Where("id = ?", id).Updates(user).Error
}

// Delete a user
func DeleteUser(id int) error {
    return config.DB.Delete(&models.User{}, id).Error
}
