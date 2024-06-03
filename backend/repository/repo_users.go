package repository

import (
	"time"

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
				Usu:         "Lau153",
				Password:    "15647k",
				Automobiles: 0,
				Types_cars:  map[string][]time.Time{},
				Debts:       0,
			},
			"Lily1526": {
				Usu:         "Lily1526",
				Password:    "1234",
				Automobiles: 1,
				Types_cars: map[string][]time.Time{
					"ref0001": {
						time.Date(2024, 06, 02, 10, 0, 0, 0, time.UTC),
						time.Date(2024, 06, 03, 10, 0, 0, 0, time.UTC),
					},
				},
				Debts: 30,
			},
		},
	}
}
