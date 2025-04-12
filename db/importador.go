package db

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func ImportarCSV() {
	file, err := os.Open("Listado_egresos_hospitalarios_ene2022_mar2025.csv")
	if err != nil {
		fmt.Println("Error abriendo CSV:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // salta la cabecera

	contador := 0

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		edad, _ := strconv.Atoi(record[4])
		fechaIngreso, _ := time.Parse("20060102", record[5])
		fechaEgreso, _ := time.Parse("20060102", record[6])
		ubigeo, _ := strconv.Atoi(record[7])
		fechaCorte, _ := time.Parse("20060102", record[9])

		_, err = DB.Exec(`
			INSERT INTO egresosh (uuid, historia, sexo, edad, fecha_ingreso, fecha_egreso, ubigeo_lugar_residencia, lugar_residencia, fecha_corte)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			record[1], record[2], record[3], edad, fechaIngreso, fechaEgreso, ubigeo, record[8], fechaCorte)

		if err != nil {
			fmt.Println("Error insertando registro:", err)
		} else {
			contador++
			fmt.Printf("Registro %d insertado correctamente\n", contador)
		}
	}

	fmt.Println("Importación completada")
}
