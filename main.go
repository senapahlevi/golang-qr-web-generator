package main

import (
	"log"
	"qrweb/customer"
	"qrweb/databases"
	"qrweb/generateqr"
	"qrweb/login"
	"qrweb/middleware"

	"qrweb/register"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := databases.SetDB()
	if err != nil {
		log.Fatal(err)
	}
	login.SetDatabaseLogin(db)
	register.SetDatabaseRegister(db)
	customer.SetDatabase(db)
	generateqr.SetQRCodeController(db)
	// router := gin.Default()
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*", "http://localhost:3000"},
	// 	AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	// AllowHeaders:     []string{"Content-Type"},
	// 	AllowHeaders:     []string{"application/json"},
	// 	ExposeHeaders:    []string{"application/x-www-form-urlencoded"},
	// 	AllowCredentials: true,
	// }))
	// api := router.Group("/api/")

	router := gin.Default()
	api := router.Group("/api")

	// Setup the CORS middleware for the api group router
	api.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		// AllowHeaders:     []string{"Content-Type"}, // or []string{"*"} to allow any header
		AllowHeaders:     []string{"*", "http://localhost:3000"}, // or []string{"*"} to allow any header
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
	}))

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "https://next-bf-home-routes.vercel.app", "*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
		// ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// api.GET("/customers", customer.GetCustomer)
	api.POST("/create-customer", customer.CreateCustomer)
	api.GET("/customer/:id", customer.GetCustomerID)
	api.PUT("/customer/:id", customer.UpdateDataCustomer)
	api.GET("/customers", middleware.AuthMiddleware(), customer.GetCustomer)

	//genereate qr code
	api.POST("/qr-generate", generateqr.GenerateQRCode)

	//login register
	api.POST("/login", login.LoginCustomer)
	api.POST("/login-google", login.LoginCustomer)
	api.POST("/register", register.RegisterCustomer)
	// api.POST("/register", register.RegisterCustomer)

	router.Run(":8080")

}

// func main() {
// 	AnagramText("cina anic")
// 	FindSmallSecond([]int{4, 2, 6, 1, 3})
// }

// func SortString(input string) []string {
// 	s := strings.Split(input, "")
// 	sort.Strings(s)
// 	return s
// }

// func AnagramText(input string) string {
// 	// hasil := strings.Split(input, " ")
// 	input1 := "nciaa"
// 	input2 := "acina"

// 	str1 := SortString(input1)
// 	str2 := SortString(input2)
// 	if len(str1) != len(str2) {
// 		fmt.Println("false")

// 		return "false"
// 	}

// 	for i := 0; i < len(str1); i++ {
// 		if str1[i] != str2[i] {
// 			fmt.Println("false")

// 			return "false"
// 		}
// 	}
// 	fmt.Println("true")
// 	return "true"
// }

// func FindSmallSecond(input []int) {
// 	sort.Sort(sort.IntSlice(input))
// 	fmt.Println("hasile", input[1])

// }

// func FrequentShow(input []int) {

// }
