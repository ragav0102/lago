package handlers

import (
	c "../controllers"
	d "../decorators"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
	var users = c.UsersIndex()

	Respond200(w, r, d.UsersIndexDecorator(users))
}

func UsersObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var user_id, _ = strconv.Atoi(vars["id"])
	var user = c.UserDetails(user_id)

	if user.ID == 0 {
		Respond404(w, r)
		return
	}

	Respond200(w, r, d.UserDetailsDecorator(user))
}
