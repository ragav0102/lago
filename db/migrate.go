package main

import (
	d "../database"
	"../models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
}

func main() {
	d.ConnectDataBase()
	d.DB.AutoMigrate(&models.User{}, &models.WorkOrder{}, &models.CreditCard{})
}
