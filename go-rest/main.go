package main

import (
	"fmt"
	"goRest/database"
	"goRest/routes"
)

func main() {
	database.DBConnect()

	fmt.Println("Iniciando servidor:")
	routes.HandleRequest()
}
