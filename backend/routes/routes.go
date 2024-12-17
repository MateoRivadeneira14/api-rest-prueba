package routes

import (
	"api-rest-prueba/backend/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/api/clientes", handlers.GetClientes)
}
