package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/controllers"
)

/*
	Funciones encargadas de atender las solicitudes relacionadas a la base de datos de los usuarios
*/

type HandlerAutomobiles struct {
	controller_auto *controllers.AutoController
}

func NewHandlerAutomobiles(controller_auto *controllers.AutoController) (*HandlerAutomobiles, error) {
	if controller_auto == nil {
		return nil, fmt.Errorf("para instanciar un handler es necesario un controllador no nulo")
	}
	return &HandlerAutomobiles{
		controller_auto: controller_auto,
	}, nil
}

func (h *HandlerAutomobiles) ListarAutos(writer http.ResponseWriter, req *http.Request) {
	/*
		Retorna toda la información de la base de datos de automobiles.
	*/
	autos, err := h.controller_auto.ListarAutos(100, 0)
	if err != nil {
		log.Printf("fallo al leer autos, con error: %s", err.Error())
		http.Error(writer, "fallo al leer los autos", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(autos)
}

func (h *HandlerAutomobiles) NuevoAuto(writer http.ResponseWriter, req *http.Request) {
	/*
		POST
	*/
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al crear un nuevo auto, con error: %s", err.Error())
		http.Error(writer, "fallo al crear un auto", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	nuevoId, err := h.controller_auto.CrearAuto(body)
	if err != nil {
		log.Println("fallo al crear un auto, con error:", err.Error())
		http.Error(writer, "fallo al crear un auto", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(nuevoId))
}

func (h *HandlerAutomobiles) TraerAutos(writer http.ResponseWriter, req *http.Request) {
	/*
		Realiza un filtrado en la base de datos por categorías como: type_transmission,
		type_fuel, year, model, color, price, seats y brand. Cuando es por price retorna
		los menores e iguales al precio elegido mientras que cuando es por asientos
		retorna los mayores e iguales.
	*/
	filter := mux.Vars(req)["filter"]
	kind := mux.Vars(req)["kind"]
	fmt.Println("FILTER: ", filter)
	fmt.Println("KIND: ", kind)
	autos, err := h.controller_auto.FiltrarAutos(filter, kind, 100, 0) ///subir es limite (1000)
	if err != nil {
		log.Printf("fallo al leer autos, con error: %s", err.Error())
		http.Error(writer, "fallo al leer los autos", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(autos))
}

func (h *HandlerAutomobiles) ActualizarAuto(writer http.ResponseWriter, req *http.Request) {
	/*
		 	Función que permite actualizar la información asociada a un auto existente en
			la base de datos de automobiles.
	*/
	ref := mux.Vars(req)["ref"]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al actualizar un auto, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un auto, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	err = h.controller_auto.ActualizarUnAuto(body, ref)
	if err != nil {
		log.Printf("fallo al actualizar un auto, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un auto, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
