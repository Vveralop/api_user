package controller

import (
	"gogin/service"
	"net/http"
	"strconv"

	"gogin/models"

	"github.com/gin-gonic/gin"
)

// Get all users
func GetUsers(c *gin.Context) {
    users, err := service.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
        return
    }
    c.JSON(http.StatusOK, users)
}

// Get user by ID
func GetUserByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    user, err := service.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

// Create a new user
func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if err := service.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
        return
    }

    c.JSON(http.StatusOK, user)
}

// Update user by ID
func UpdateUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if err := service.UpdateUser(id, &user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
        return
    }

    c.JSON(http.StatusOK, user)
}

// Delete user by ID
func DeleteUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    if err := service.DeleteUser(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
