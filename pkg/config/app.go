package config

//* This MySql import is different from the imports in mysql-grpc file
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Purpose of this file is to return a variable called db
// Sothat other files can talk to the database

var (db * gorm.DB)

func Connect(){
	// helps connect with sql database
	d, err := gorm.Open("mysql", "root:Chowdhury0511@@tcp(127.0.0.1:3306)/vocab?charset=utf8&parseTime=True&loc=Local")
	// id: root, password: Chowdhury0511@, database-name: vocab
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB()*gorm.DB{
	return db
}