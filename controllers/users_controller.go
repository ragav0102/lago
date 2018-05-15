package controllers

import (
	d "../database"
	m "../models"
)

func UsersIndex() []m.User {
	var users = []m.User{}

	d.DB.Find(&users)

	return users
}

func UserDetails(userId int) m.User {
	var user m.User

	d.DB.Preload("CreditCard").First(&user, userId)

	return user
}
