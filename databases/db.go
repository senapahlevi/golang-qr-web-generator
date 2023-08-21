package databases

import (
	"log"
	"qrweb/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type Database struct {
// 	DB *gorm.DB
// }

// func SetDB() (*Database, error) {

// 	host := os.Getenv("DB_HOST")
// 	user := os.Getenv("DB_USER")
// 	dbnames := os.Getenv("DB_NAME")
// 	pass := os.Getenv("DB_PASS")
// 	dbport := os.Getenv("DB_PORT")

// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Jakarta", host, user, pass, dbnames, dbport)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// errs := db.AutoMigrate(&models.User{})
// 	// if errs != nil {
// 	// 	log.Fatal(errs)
// 	// }

// 	return &Database{
// 		DB: db,
// 	}, nil
// }

type DB struct {
	DB *gorm.DB
}

func SetDB() (*DB, error) {
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Jakarta", host, user, pass, dbnames, dbport)

	dbsetting := "root:@tcp(localhost:3306)/qrweb?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbsetting), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = db.AutoMigrate(&models.Customer{}, &models.RolesUser{})
	return &DB{
		DB: db,
	}, nil
}
