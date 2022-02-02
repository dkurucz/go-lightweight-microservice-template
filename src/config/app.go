package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"fmt"
)

var (
	db * gorm.DB
)

func Connect(){
	MYSQL_USER, ok := os.LookupEnv("MYSQL_USER")
	if !ok {
		fmt.Printf("Err: the env var MYSQL_USER is not set")
		return
	}
	MYSQL_PASSWORD, ok := os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		fmt.Printf("Err: the env var MYSQL_PASSWORD is not set")
		return
	}
	MYSQL_DATABASE, ok := os.LookupEnv("MYSQL_DATABASE")
	if !ok {
		fmt.Printf("Err: the env var MYSQL_DATABASE is not set")
		return
	}
	MYSQL_HOST, ok := os.LookupEnv("MYSQL_HOST")
	if !ok {
		fmt.Printf("Err: the env var MYSQL_HOST is not set")
		return
	}
	MYSQL_PORT, ok := os.LookupEnv("MYSQL_PORT")
	if !ok {
		fmt.Printf("Err: the env var MYSQL_PORT is not set")
		return
	}

	d, err := gorm.Open("mysql", "" + MYSQL_USER + ":" + MYSQL_PASSWORD + "@tcp(" + MYSQL_HOST + ":" + MYSQL_PORT + ")" + "/" + MYSQL_DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}