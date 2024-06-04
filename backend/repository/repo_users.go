package repository

import (
	models "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/models"
)

// Persistencia en memoria:
type BaseDatosUsers struct {
	Memoria map[string]models.User
}

// Función para crear un nuevo repositorio:
// con esto se devuelve base de datos inicializada
func NewBaseDatosUsers() *BaseDatosUsers {
	return &BaseDatosUsers{
		//Memoria: make(map[string]models.User),
		Memoria: map[string]models.User{
			"Lau153": {
				Usu:      "Lau153",
				Name:     "Laura Gomez",
				Email:    "lau.gomez@gmail.com",
				Password: "15647k",
			},
			"Lily1526": {
				Usu:      "Lily1526",
				Name:     "Lily Vetancur",
				Email:    "lily_vetancur@hotmail.com",
				Password: "1234",
			},
		},
	}
}
