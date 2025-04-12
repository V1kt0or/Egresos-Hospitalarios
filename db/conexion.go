package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Conectar() {
	var err error
	DB, err = sql.Open("sqlite3", "./egresosh.db")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a SQLite:", err)
	}

	log.Println("Conexión exitosa a SQLite")

	crearTablaSiNoExiste()
}

func crearTablaSiNoExiste() {
	query := `
	CREATE TABLE IF NOT EXISTS egresosh (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid TEXT,
		historia TEXT,
		sexo TEXT,
		edad INTEGER,
		fecha_ingreso TEXT,
		fecha_egreso TEXT,
		ubigeo_lugar_residencia INTEGER,
		lugar_residencia TEXT,
		fecha_corte TEXT
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creando la tabla: %v", err)
	}
}
