package routers

import (
	h "../handlers"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRouter() *mux.Router {
	fmt.Println("Setting up router...")
	r := mux.NewRouter()
	addRoutes(r)
	return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal("Alive")
	h.Respond200(w, r, b)
}

func addRoutes(r *mux.Router) {
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/users", h.UsersIndex)
	r.HandleFunc("/users/{id:[0-9]+}", h.UsersObject)
}
