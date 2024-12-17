package handlers

import (
	"api-rest-prueba/backend/db"
	"api-rest-prueba/backend/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Obtener todos los clientes
func GetClientes(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, nombre, correo_electronico, numero_telefono FROM clientes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var clientes []models.Cliente
	for rows.Next() {
		var cliente models.Cliente
		if err := rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.CorreoElectronico, &cliente.NumeroTelefono); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clientes = append(clientes, cliente)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

// Crear un cliente
func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var cliente models.Cliente
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cliente); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insertar en la base de datos
	_, err := db.DB.Exec("INSERT INTO clientes (nombre, correo_electronico, numero_telefono) VALUES (?, ?, ?)", cliente.Nombre, cliente.CorreoElectronico, cliente.NumeroTelefono)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Cliente creado")
}

// Actualizar un cliente
func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	var cliente models.Cliente
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cliente); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]
	_, err := db.DB.Exec("UPDATE clientes SET nombre = ?, correo_electronico = ?, numero_telefono = ? WHERE id = ?", cliente.Nombre, cliente.CorreoElectronico, cliente.NumeroTelefono, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Cliente actualizado")
}

// Eliminar un cliente
func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := db.DB.Exec("DELETE FROM clientes WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
