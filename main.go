package main

import (
	"log"
	"qrweb/customer"
	"qrweb/databases"

	"github.com/gin-gonic/gin"
)

// import (
// 	"log"
// 	"qrweb/algobash"
// 	_ "qrweb/algobash"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// )

// // func init() {

// // 	err := godotenv.Load(".env")

// // 	if err != nil {
// // 		log.Fatal("Error loading .env file")
// // 	}
// // }

// // // func Duplicate(coba []int) {
// // // 	cobas := []string{"123"}
// // // 	ccooba = cobas.Split(cobas, "")
// // // 	for i := 0; i < len(coba)-1; i++ {
// // // 		if coba[i] != coba[i+1] {
// // // 			// if int.Contains(cobas,i){
// // // 			// 	cobas = append(cobas, coba[i+1])
// // // 			// }
// // // 			// if (slices.Contains(cobas, "foo"))
// // // 			// fmt.Println(cobas[i])
// // // 		}
// // // 		// for j := 0; j < len(coba)-1; j++ {
// // // 		// }
// // // 	}
// // // 	fmt.Println(cobas, "hello")
// // // 	fmt.Println(strings.Contains(cobas, "123"))
// // // }

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
