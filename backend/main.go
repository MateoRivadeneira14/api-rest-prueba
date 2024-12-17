package main

import (
	"api-rest-prueba/backend/db"
	"api-rest-prueba/backend/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Conexión a la base de datos
	db.Connect()

	// Configuración del router
	r := mux.NewRouter()

	// Rutas API
	r.HandleFunc("/clientes", handlers.GetClientes).Methods("GET")
	r.HandleFunc("/clientes", handlers.CreateCliente).Methods("POST")
	r.HandleFunc("/clientes/{id}", handlers.UpdateCliente).Methods("PUT")
	r.HandleFunc("/clientes/{id}", handlers.DeleteCliente).Methods("DELETE")

	// Ruta para archivos estáticos (frontend)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/static"))))

	// Servidor
	fmt.Println("Servidor iniciado en el puerto 8080")
	http.ListenAndServe(":8080", r)
}
