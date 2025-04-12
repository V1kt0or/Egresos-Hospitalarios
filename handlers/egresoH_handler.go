package handlers

import (
	"MS1_Egresos_Hospitalarios/db"
	"MS1_Egresos_Hospitalarios/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux" // Importa mux para manejar rutas con parámetros
)

func CrearObjeto(w http.ResponseWriter, r *http.Request) {
	var obj models.Egresos
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	err := db.DB.QueryRow(
		"INSERT INTO egresosh (uuid, historia, sexo, edad, fecha_ingreso, fecha_egreso, ubigeo_lugar_residencia, lugar_residencia, fecha_corte) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		obj.UUID, obj.Historia, obj.Sexo, obj.Edad, obj.FechaIngreso, obj.FechaEgreso, obj.UbigeoLugarResidencia, obj.LugarResidencia, obj.FechaCorte,
	).Scan(&obj.ID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error al insertar: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(obj)
}

func VerObjetos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, uuid, historia, sexo, edad, fecha_ingreso, fecha_egreso, ubigeo_lugar_residencia, lugar_residencia, fecha_corte FROM egresosh")
	if err != nil {
		http.Error(w, "Error al consultar", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var lista []models.Egresos
	for rows.Next() {
		var obj models.Egresos
		err := rows.Scan(&obj.ID, &obj.UUID, &obj.Historia, &obj.Sexo, &obj.Edad, &obj.FechaIngreso, &obj.FechaEgreso, &obj.UbigeoLugarResidencia, &obj.LugarResidencia, &obj.FechaCorte)
		if err != nil {
			http.Error(w, "Error al leer fila", http.StatusInternalServerError)
			return
		}
		lista = append(lista, obj)
	}

	json.NewEncoder(w).Encode(lista)
}

// Función para obtener un solo objeto por su ID
func VerObjetoPorID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var (
		fechaIngresoStr string
		fechaEgresoStr  string
		fechaCorteStr   string
		obj             models.Egresos
	)

	err := db.DB.QueryRow(
		`SELECT id, uuid, historia, sexo, edad, fecha_ingreso, fecha_egreso, ubigeo_lugar_residencia, lugar_residencia, fecha_corte 
		FROM egresosh WHERE id = $1`, id,
	).Scan(
		&obj.ID, &obj.UUID, &obj.Historia, &obj.Sexo, &obj.Edad,
		&fechaIngresoStr, &fechaEgresoStr, &obj.UbigeoLugarResidencia, &obj.LugarResidencia, &fechaCorteStr,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No se encontró el objeto con ese ID", http.StatusNotFound)
		} else {
			http.Error(w, fmt.Sprintf("Error al consultar: %v", err), http.StatusInternalServerError)
		}
		return
	}

	// Parsear fechas
	obj.FechaIngreso, _ = time.Parse("2006-01-02 15:04:05", fechaIngresoStr)
	obj.FechaEgreso, _ = time.Parse("2006-01-02 15:04:05", fechaEgresoStr)
	obj.FechaCorte, _ = time.Parse("2006-01-02 15:04:05", fechaCorteStr)

	json.NewEncoder(w).Encode(obj)
}
