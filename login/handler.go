package login

import (
	"net/http"
	"qrweb/databases"
	"qrweb/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var db *gorm.DB

var secret_key = "tokentoken"

func SetDatabaseLogin(database *databases.DB) {
	db = database.DB
}

func LoginCustomer(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	var customer models.Customer
	if err := db.Where("username = ?", request.Username).First(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid username or password",
			"status": "Denied",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid username or password",
			"status": "Denied",
		})
		return
	}
	token, err := generateToken(customer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"status": "success",
	})
}

func generateToken(CustomerID int) (string, error) {
	claims := jwt.MapClaims{
		"id":  CustomerID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return "", err
	}
	return signedToken, err
}

func LoginGoogleCustomer(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	var customer models.Customer
	if err := db.Where("username = ?", request.Username).First(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid username or password",
			"status": "Denied",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid username or password",
			"status": "Denied",
		})
		return
	}
	token, err := generateToken(customer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"status": "success",
	})
}
