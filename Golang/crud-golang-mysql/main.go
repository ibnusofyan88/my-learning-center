package main

import (
	"net/http"

	"github.com/ibnusofyan88/crud-golang-mysql/database"
	"github.com/ibnusofyan88/crud-golang-mysql/routes"
)

func main() {
	database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server)

	http.ListenAndServe(":8080", server)
}
