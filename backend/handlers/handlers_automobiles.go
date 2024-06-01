package handlers

import (
	"encoding/json"
	"net/http"

	models "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/models"
	repository "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/repository"
)

/*
	Funciones encargadas de implementar las funciones relacionadas a la base de datos de los autos
*/

type HandlerAutos struct {
	BD *repository.BaseDatosAutomobiles
}

func NewHandlerAutos(bd *repository.BaseDatosAutomobiles) *HandlerAutos {
	return &HandlerAutos{
		BD: bd,
	}
}

func (hc *HandlerAutos) ListarAutos() http.HandlerFunc {
	/*
		Función que retorna toda la información de la base de datos de autos.
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		autos := []models.Automobile{}
		for _, auto := range hc.BD.Memoria {
			autos = append(autos, auto)
		}
		jsonCars, err := json.Marshal(autos)
		if err != nil {
			http.Error(w, "fallo al comunicar en json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonCars)
	})
}

/*
func (hc *HandlerAutos) TraerAutos() http.HandlerFunc {
	/*
		Función que retorna todos los autos disponibles bajo cierta categoría.
		La categoría puede ser:
			Type_transmission string    //puede ser: manual ó automático
			Type_fuel         string    //puede ser: gasolina, Diesel ó eléctrico
			Year              int       //año de fabricación del modelo del carro
			Model             string    //modelo del carro en stock
			Color             string    //color del carro en stock
			Price             int       //costo de renta del carro por día
			Seats             int       //capacidad de personas que soporta el carro
			Brand             string    //marca del carro en stock
	///////////
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		autos := []models.Automobile{}
		for _, auto := range hc.BD.Memoria {
			autos = append(autos, auto)
		}
		jsonCars, err := json.Marshal(autos)
		if err != nil {
			http.Error(w, "fallo al comunicar en json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonCars)
	})
}*/
