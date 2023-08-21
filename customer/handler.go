package customer

import (
	"fmt"
	"net/http"
	"qrweb/databases"
	"qrweb/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDatabase(database *databases.DB) {
	db = database.DB
}
func GetCustomer(c *gin.Context) {
	customer := []models.Customer{}

	err := db.Limit(5).Find(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
	}
	c.JSON(http.StatusOK, customer)
	return
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	fmt.Println("hello")
	result := db.Create(&customer)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})

	}

	c.JSON(http.StatusOK, customer)
}

func GetCustomerID(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")
	result := db.Where("id = ?", id).First(&customer)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}
	c.JSON(http.StatusOK, customer)
}

func UpdateDataCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")
	result := db.Model(&customer).Where("id = ?", id).First(&customer)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result = db.Save(&customer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})

	}
	c.JSON(http.StatusOK, customer)
}

func DeleteDataCustomer(c *gin.Context) {

	// var house House
	var customer models.Customer
	id := c.Param("id")
	result := db.Where("id = ?", id).First(&customer)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data customer not found"})
		return
	}

	result = db.Delete(&customer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)

}
