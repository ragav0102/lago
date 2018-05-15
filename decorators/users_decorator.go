package decorators

import (
	m "../models"
	"encoding/json"
)

func UserDetailsDecorator(user m.User) []byte {
	item := user.UserObject()

	item["credit_card_no"] = user.CreditCard.Number

	resp, _ := json.Marshal(item)
	return resp
}

func UsersIndexDecorator(users []m.User) []byte {
	items := make([]interface{}, 0)
	for _, item := range users {
		items = append(items, item.UserObject())
	}
	resp, _ := json.Marshal(items)
	return resp
}
