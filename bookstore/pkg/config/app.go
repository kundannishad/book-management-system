package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

const (
	username = "root"
	password = "Password"
	hostname = "127.0.0.1:3306"
	databse  = "golang_bookstore"
)

func Dsn(dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, databse)
}

func Connect() {

	fmt.Println(Dsn(""))

	//d, err := gorm.Open("mysql", "root:Password/simplerest?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", Dsn("")+"?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
