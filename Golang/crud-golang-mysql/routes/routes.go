package routes

import (
	"net/http"

	"github.com/ibnusofyan88/crud-golang-mysql/controller"
)

func MapRoutes(server *http.ServeMux) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/employee", controller.NewHIndexEmployeeController())
}
