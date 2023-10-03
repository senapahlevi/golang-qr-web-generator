package register

import (
	"net/http"
	"qrweb/databases"
	"qrweb/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// import (
// 	"gorm.io/gorm"
// )

// var db *gorm.DB
// var secret_key = "tokentoken"

// func SetDatabaseRegister(databased *databases) {
// 	// db = databased.DB
// }
var db *gorm.DB

var secret_key = "tokentoken"

func SetDatabaseRegister(databases *databases.DB) {
	db = databases.DB
}

func RegisterCustomer(c *gin.Context) {
	var register RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//check email
	var existEmail models.Customer
	if err := db.Where("email = ?", register.Email).First(&existEmail).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already taken ", "status": "denied"})
		return
	}

	//check username
	var existUsername models.Customer
	if err := db.Where("username = ?", register.Username).First(&existUsername).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already taken ", "status": "denied"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password", "status": "denied"})
		return
	}
	customer := models.Customer{
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Username:  register.Username,
		Email:     register.Email,
		Password:  string(hashedPassword),
	}

	if err := db.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "status": "denied"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})

	return
}
