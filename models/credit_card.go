package models

import (
	"github.com/jinzhu/gorm"
)

type CreditCard struct {
	gorm.Model
	Number      string
	ExpiryMonth uint
	ExpiryYear  uint
	CVV         uint
}
