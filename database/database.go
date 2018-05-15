package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	// DBCon is the connection handle
	// for the database
	DB *gorm.DB
)

func ConnectDataBase() {
	// Initiates one connection to DB and holds it until program terminates.
	// Extra required connections will be handled inrernally
	DB, _ = gorm.Open("mysql", "root:@/freshfield1?charset=utf8&parseTime=True&loc=Local")
	DB.LogMode(true)
	// defer DB.Close()
}
