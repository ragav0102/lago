package main

import (
	d "../database"
	"../models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	d.ConnectDataBase()
	d.DB.Create(&models.User{FirstName: "John", LastName: "Wick", Email: "johnwick@johnwick.com", CreditCard: models.CreditCard{Number: "312442141241", ExpiryMonth: 9, ExpiryYear: 2022, CVV: 487}})
	d.DB.Create(&models.User{FirstName: "Rich", LastName: "Mike", Email: "richmike@richmike.com", CreditCard: models.CreditCard{Number: "632313123121", ExpiryMonth: 11, ExpiryYear: 2020, CVV: 212}})
	d.DB.Create(&models.WorkOrder{ExternalId: 234, OrderType: "dish setup", ContactId: 544})
}
