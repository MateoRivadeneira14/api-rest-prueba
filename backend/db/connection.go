package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	// Cambiar por tu usuario y contraseña de MariaDB
	DB, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/tienda")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}
	fmt.Println("Conexión a la base de datos establecida.")
}
