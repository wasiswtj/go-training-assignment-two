package config

import (
	"assignment-two/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "assignment-two"
)

func ConnectGorm() *gorm.DB {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	CheckError(err)

	// check db
	// err = db.Ping()
	CheckError(err)

	fmt.Println("Connected Gorm Sql!")

	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.Item{})
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
