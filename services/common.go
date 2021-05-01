package services

import (
	"github.com/jinzhu/gorm"
)

var username = "test"
var password = "passw0rd"
var dbHost = "database"
var dbPort = "5432"
var dbName = "apidb"

var Db *gorm.DB

func OpenDatabase() {
	//open a db connection
	var err error
	Db, err = gorm.Open("postgres", "postgres://"+username+":"+password+"@"+dbHost+":"+dbPort+"/"+dbName+"?sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
}
