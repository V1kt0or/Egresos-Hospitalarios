package main

import (
	"MS1_Egresos_Hospitalarios/db"
	"MS1_Egresos_Hospitalarios/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux" // Importa mux para manejar rutas con parámetros
)

func main() {
	db.Conectar()
	db.ImportarCSV()

	// Crear un nuevo enrutador
	r := mux.NewRouter()

	// Definir las rutas
	r.HandleFunc("/crear", handlers.CrearObjeto).Methods("POST")
	r.HandleFunc("/ver", handlers.VerObjetos).Methods("GET")
	r.HandleFunc("/ver/{id}", handlers.VerObjetoPorID).Methods("GET") // Ruta para obtener un objeto por su ID

	// Configurar el servidor HTTP
	log.Println("Servidor en http://localhost:8191")
	log.Fatal(http.ListenAndServe(":8191", r)) // Usar el enrutador mux
}
