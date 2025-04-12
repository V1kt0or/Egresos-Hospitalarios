package models

import "time"

type Egresos struct {
	ID                    int       `json:"id"`
	UUID                  string    `json:"uuid"`
	Historia              string    `json:"historia"`
	Sexo                  string    `json:"sexo"`
	Edad                  int       `json:"edad"`
	FechaIngreso          time.Time `json:"fecha_ingreso"`
	FechaEgreso           time.Time `json:"fecha_egreso"`
	UbigeoLugarResidencia int       `json:"ubigeo_lugar_residencia"`
	LugarResidencia       string    `json:"lugar_residencia"`
	FechaCorte            time.Time `json:"fecha_corte"`
}
