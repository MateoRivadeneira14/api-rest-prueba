package models

type Cliente struct {
	ID                int    `json:"id"`
	Nombre            string `json:"nombre"`
	CorreoElectronico string `json:"correo_electronico"`
	NumeroTelefono    string `json:"numero_telefono"`
}
