package repository

import (
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
				Year:              2024,
				Model:             "mazda3",
				Color:             "rojo",
				Price:             30,
				Seats:             5,
				Brand:             "mazda",
				Image:             "https://www.mazda.com.co/globalassets/cars/mazda-3-2024/versiones-y-menu/desplegable-vehiculos.png",
				Quantity:          20,
			},
			"ref0002": {
				Ref:               "ref0002",
				Type_transmission: "manual",
				Type_fuel:         "gasolina",
				Year:              2023,
				Model:             "sandero life+",
				Color:             "azul",
				Price:             20,
				Seats:             5,
				Brand:             "renault",
				Image:             "https://acroadtrip.blob.core.windows.net/catalogo-imagenes/l/RT_V_9e1b416d37674588b9959680073c1e92.webp",
				Quantity:          10,
			},
			"ref0003": {
				Ref:               "ref0003",
				Type_transmission: "automatico",
				Type_fuel:         "gasolina",
				Year:              2010,
				Model:             "escape 3.0",
				Color:             "negro",
				Price:             16,
				Seats:             5,
				Brand:             "ford",
				Image:             "https://www.vehiclehistory.com/uploads/2010-Ford-Escape.jpg",
				Quantity:          15,
			},
			"ref0004": {
				Ref:               "ref0004",
				Type_transmission: "automatico",
				Type_fuel:         "gasolina",
				Year:              2010,
				Model:             "journey 2.4",
				Color:             "rojo",
				Price:             20,
				Seats:             7,
				Brand:             "dodge",
				Image:             "https://media.ed.edmunds-media.com/dodge/journey/2010/oem/2010_dodge_journey_4dr-suv_rt_fq_oem_2_1600.jpg",
				Quantity:          5,
			},
		},
	}
}
