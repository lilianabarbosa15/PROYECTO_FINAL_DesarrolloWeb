package repository

import (
	"time"

	models "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/models"
)

// Persistencia en memoria:
type BaseDatosAutomobiles struct {
	Memoria map[string]models.Automobile
}

// Función para crear un nuevo repositorio:
// con esto se devuelve base de datos inicializada
func NewBaseDatosAutomobiles() *BaseDatosAutomobiles {
	return &BaseDatosAutomobiles{
		Memoria: map[string]models.Automobile{
			"ref0001": {
				Ref:               "ref0001",
				Type_transmission: "automatico",
				Type_fuel:         "gasolina",
				Model:             "mazda3",
				Color:             "rojo",
				Price:             30,
				Seats:             5,
				Brand:             "mazda",
				Image:             "https://www.mazda.com.co/globalassets/cars/mazda-3-2024/versiones-y-menu/desplegable-vehiculos.png ",
				Daysinuse:         1,
				Usedby:            "Lily1526",
				Available:         time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 10, 0, 0, 0, time.UTC),
			},
		},
	}
}
