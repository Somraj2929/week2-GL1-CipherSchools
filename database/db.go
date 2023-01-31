package database

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/somraj2929/week2-GL1-CipherSchools/models"
	//"github.com/jinzhu/gorm"
)


var DB *gorm.DB


func GetDB() *gorm.DB {
	return DB
}

func Setup(){
	host:= "localhost"
	port:= "5432"
	dbName:= "book"
	username:= "postgres"
	password:= "Rock2929"
	arg := "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbName + " sslmode=disable password="+ password
	db, err := gorm.Open(postgres.Open(arg), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB = db
}