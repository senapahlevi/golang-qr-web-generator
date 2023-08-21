package algobash

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	// err = db.AutoMigrate(&Customer{})
	return &DB{
		DB: db,
	}, nil
}
