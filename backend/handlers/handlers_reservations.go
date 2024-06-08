package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/controllers"
)

/*
	Funciones encargadas de atender las solicitudes relacionadas a la base de datos de los usuarios
*/

type HandlerReservas struct {
	controller *controllers.ReservaController
}

func NewHandlerReservas(controller *controllers.ReservaController) (*HandlerReservas, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler es necesario un controllador no nulo")
	}
	return &HandlerReservas{
		controller: controller,
	}, nil
}

func (h *HandlerReservas) ListarReservas(writer http.ResponseWriter, req *http.Request) {
	/*
		Retorna toda la información de la base de datos de las reservas.
	*/
	reservas, err := h.controller.ListarReservas(100, 0)
	if err != nil {
		log.Printf("fallo al leer reserva, con error: %s", err.Error())
		http.Error(writer, "fallo al leer las reservas", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(reservas)
}

func (h *HandlerReservas) NuevaReserva(writer http.ResponseWriter, req *http.Request) {
	/*
		Crea las nuevas reservas través de POST.
	*/
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al crear nueva reserva, con error: %s", err.Error())
		http.Error(writer, "fallo al crear nueva reserva", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	nuevoId, err := h.controller.CrearReserva(body)
	if err != nil {
		log.Println("fallo al crear nueva reserva, con error:", err.Error())
		http.Error(writer, "fallo al crear nueva reserva", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("id nueva reserva: %d", nuevoId)))
}

func (h *HandlerReservas) BorrarReserva(writer http.ResponseWriter, req *http.Request) {
	/*
		Función que se encarga de borrar una reserva existente, asociada a cierto
		usuario en especifico.
	*/
	usu := mux.Vars(req)["usu"]
	ref := mux.Vars(req)["ref"]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al borrar reserva, con error: %s", err.Error())
		http.Error(writer, "fallo al borrar reserva", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	h.controller.BorrarUnaReserva(body, usu, ref)
	writer.WriteHeader(http.StatusOK)
}

func (h *HandlerReservas) ActualizarReserva(writer http.ResponseWriter, req *http.Request) {
	/*
		Función que se encarga de actualizar el estado de los detalles de
		cierto de la reserva de cierto usuario en especifico.
	*/
	usu := mux.Vars(req)["usu"]
	ref := mux.Vars(req)["ref"]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al actualizar reserva, con error: %s", err.Error())
		http.Error(writer, "fallo al actualizar reserva", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	h.controller.ActualizarUnaReserva(body, usu, ref)
	writer.WriteHeader(http.StatusOK)
}
