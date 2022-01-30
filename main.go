package main

import (
	"alura/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8900", nil)
}
