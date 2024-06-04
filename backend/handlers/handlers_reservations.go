package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/models"
	repository "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/repository"
)

/*
	Funciones encargadas de implementar las funciones relacionadas a la base de datos de las reservas
*/

type HandlerReservas struct {
	BD *repository.BaseDatosReservas
}

func NewHandlerReservas(bd *repository.BaseDatosReservas) *HandlerReservas {
	return &HandlerReservas{
		BD: bd,
	}
}

func (hc *HandlerReservas) ListarReservas() http.HandlerFunc {
	/*
		Función que retorna toda la información de la base de datos de usuarios.
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reservas := []models.Reserva{}
		for _, reserva := range hc.BD.Memoria {
			reservas = append(reservas, reserva)
		}
		jsonUsu, err := json.Marshal(reservas)
		if err != nil {
			http.Error(w, "fallo al comunicar en json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonUsu)
	})
}

func (hc *HandlerReservas) NuevaReserva() http.HandlerFunc {
	/*
		Función que se encarga de crear una nueva reserva, asociada a cierto
		usuario en especifico.
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "fallo la peticion POST", http.StatusBadRequest)
		}
		reserva := models.Reserva{}
		err = json.Unmarshal(body, &reserva)
		if err != nil {
			http.Error(w, "fallo al codificar en json", http.StatusInternalServerError)
		}
		hc.BD.Memoria[reserva.IdUser] = reserva
		w.WriteHeader(http.StatusCreated)
	})
}

func (hc *HandlerReservas) BorrarReserva() http.HandlerFunc {
	/*
		Función que se encarga de borrar una reserva existente, asociada a cierto
		usuario en especifico.
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "fallo la peticion POST", http.StatusBadRequest)
		}
		reserva := models.Reserva{}
		err = json.Unmarshal(body, &reserva)
		if err != nil {
			http.Error(w, "fallo al codificar en json", http.StatusInternalServerError)
		}
		delete(hc.BD.Memoria, reserva.IdUser)
		w.WriteHeader(http.StatusCreated)
	})
}

func (hc *HandlerReservas) ActualizarReserva() http.HandlerFunc {
	/*
		Función que se encarga de actualizar el estado de los detalles de
		cierto de la reserva de cierto usuario en especifico.
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		change := mux.Vars(r)["change"]
		body, err := io.ReadAll(r.Body)
		fmt.Println("body: ", body)
		if err != nil {
			http.Error(w, "fallo en la peticion", http.StatusBadRequest)
		}
		reserva := models.Reserva{}
		err = json.Unmarshal(body, &reserva)
		fmt.Println("reserva: ", reserva)
		if err != nil {
			http.Error(w, "fallo al codificar en json", http.StatusInternalServerError)

		}
		reservaBase := hc.BD.Memoria[reserva.IdUser]

		if change == "add" {
			for k, v := range reserva.Details {
				reservaBase.Details[k] = v
				fmt.Println("key: ", k)
				fmt.Println("value: ", v)
				hc.BD.Memoria[reserva.IdUser] = reservaBase
			}
		} else if change == "delete" {
			for k, _ := range reserva.Details {
				delete(reservaBase.Details, k)
			}
			hc.BD.Memoria[reserva.IdUser] = reservaBase
		}

		w.WriteHeader(http.StatusCreated)
	})
}
