package main

import (
	d "./database"
	r "./routers"
	"fmt"
	"net/http"
)

func main() {
	d.ConnectDataBase()
	routeHandler := r.RegisterRouter()
	http.Handle("/", routeHandler)
	fmt.Println("Starting your server at :8080..")
	http.ListenAndServe(":8080", routeHandler)
}
