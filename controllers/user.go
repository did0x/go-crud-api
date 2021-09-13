
package controllers

import (
	models "go-crud-postgres/models"
	"go-crud-postgres/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// POST /register
// Register new user
func RegisterUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.RegisterUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create User
	user := models.User{Username: input.Username, Password: input.Password, Phone: input.Phone}
	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// POST /login/
// Login a user
func LoginUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.LoginUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var user models.User
	if err := db.Where("username = ? AND password = ?", input.Username, input.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login credentials. Please try again"})
		return
	}

	token := middleware.GetToken(input.Username)
	var resp = map[string]interface{}{
		"token": token,
		"user": input,
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})
}
