package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string
	CreditCardID uint
	CreditCard   CreditCard
}

func (u *User) FullName() string {
	return (u.FirstName + " " + u.LastName)
}

func (u *User) UserObject() map[string]interface{} {
	object := make(map[string]interface{})

	object["id"] = u.ID
	object["name"] = u.FullName()
	object["email"] = u.Email
	object["created_at"] = u.CreatedAt

	return object
}
