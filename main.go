package main

import (
	"log"
	"qrweb/customer"
	"qrweb/databases"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := databases.SetDB()
	if err != nil {
		log.Fatal(err)
	}

	customer.SetDatabase(db)
	router := gin.Default()
	api := router.Group("/api/")
	api.GET("/customers", customer.GetCustomer)
	api.POST("/create-customer", customer.CreateCustomer)
	api.GET("/customer/:id", customer.GetCustomerID)
	api.PUT("/customer/:id", customer.UpdateDataCustomer)

	router.Run(":8080")

}
