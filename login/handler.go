package login

import (
	"fmt"
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
	//original
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	var customer models.Customer
	if err := db.Where("username = ?", request.Username).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid username or password",
			"status": "Denied",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(request.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid username or password",
			"errors": err.Error(),
			"status": "Denied",
		})
		return
	}

	token, err := generateToken(customer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	fmt.Println(token, "hello")
	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"status": "success",
	})
}

// func LoginCustomer(ctx *gin.Context) {
//jika menggunakan postform
// 	// Get the username and password from the request body
// 	username := ctx.PostForm("username")
// 	password := ctx.PostForm("password")

// 	// Find the user in the database
// 	var user models.Customer
// 	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
// 		return
// 	}

// 	// Compare the password
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
// 		return
// 	}

// 	// Generate a new token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"id":  user.ID,
// 		"exp": time.Now().Add(time.Hour * 24).Unix(),
// 	})

// 	signedToken, err := token.SignedString([]byte(secret_key))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
// 		return
// 	}

// 	// Set the token in the response cookie
// 	ctx.SetCookie("token", signedToken, 60*60*24, "/", "", false, true)
// 	// coba := ctx.SetCookie("token", signedToken, 60*60*24, "/", "", false, true)
// 	// fmt.Println("hello", coba)

// 	// Respond with a success message
// 	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful"})
// }

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
