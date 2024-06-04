package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (hc *HandlerAutos) TraerAutos() http.HandlerFunc {
	/*
		Función que retorna todos los autos disponibles bajo cierta categoría.
		La categoría puede ser: type_transmission (manual ó automatico),
		type_fuel (gasolina, diesel ó electrico), year (año de fabricación del
		modelo del carro), model (modelo del carro en stock), color (color del
		carro en stock), price (costo de renta del carro por día), seats
		(capacidad de personas que soporta el carro), brand (marca del carro en
		stock)
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filter := mux.Vars(r)["filter"]
		kind := mux.Vars(r)["kind"]
		kind_int, _ := strconv.Atoi(kind)
		autos := []models.Automobile{}
		for _, auto := range hc.BD.Memoria {
			if filter == "type_transmission" {
				if auto.Type_transmission == kind {
					autos = append(autos, auto)
				}
			} else if filter == "type_fuel" {
				if auto.Type_fuel == kind {
					autos = append(autos, auto)
				}
			} else if filter == "year" {
				if strconv.Itoa(auto.Year) == kind {
					autos = append(autos, auto)
				}
			} else if filter == "model" {
				if auto.Model == kind {
					autos = append(autos, auto)
				}
			} else if filter == "color" {
				if auto.Color == kind {
					autos = append(autos, auto)
				}
			} else if filter == "price" {
				if auto.Price <= kind_int {
					autos = append(autos, auto)
				}
			} else if filter == "seats" {
				if auto.Seats >= kind_int {
					autos = append(autos, auto)
				}
			} else if filter == "brand" {
				if auto.Brand == kind {
					autos = append(autos, auto)
				}
			}
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

func (hc *HandlerAutos) NuevoAuto() http.HandlerFunc {
	/*
		Función de registro de stock, permite crear autos en la base de
		datos de automobiles.
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "fallo en la peticion", http.StatusBadRequest)
		}
		auto := models.Automobile{}
		err = json.Unmarshal(body, &auto)
		if err != nil {
			http.Error(w, "fallo al codificar en json", http.StatusInternalServerError)
		}
		hc.BD.Memoria[auto.Ref] = auto
		w.WriteHeader(http.StatusCreated)
	})
}

func (hc *HandlerAutos) ActualizarAuto() http.HandlerFunc {
	/*
		Función de actualización de cantidad de autos disponibles en stock.
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ref := mux.Vars(r)["ref"]
		if ref == "" {
			http.Error(w, "referencia no valida", http.StatusBadRequest)
			return
		}
		infoCar, ok := hc.BD.Memoria[ref]
		if !ok {
			http.Error(w, "no se encuentra informacion para esa referencia", http.StatusNotFound)
			return
		}
		quantity, _ := strconv.Atoi(mux.Vars(r)["quantity"])
		infoCar.Quantity = quantity
		hc.BD.Memoria[infoCar.Ref] = infoCar
		w.WriteHeader(http.StatusCreated)
	})
}
