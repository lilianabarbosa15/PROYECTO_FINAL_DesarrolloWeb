package handlers

import (
	"fmt"
	"io"
	_ "io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/controllers"
)

/*
	Funciones encargadas de atender las solicitudes relacionadas a la base de datos de los usuarios
*/

type HandlerUsuarios struct {
	controller *controllers.UserController
}

func NewHandlerUsuarios(controller *controllers.UserController) (*HandlerUsuarios, error) {
	if controller == nil {
		return nil, fmt.Errorf("para instanciar un handler es necesario un controllador no nulo")
	}
	return &HandlerUsuarios{
		controller: controller,
	}, nil
}

func (h *HandlerUsuarios) ListarUsuarios(writer http.ResponseWriter, req *http.Request) { // http.HandlerFunc {
	/*
		Retorna toda la información de la base de datos de usuarios.
	*/
	usuarios, err := h.controller.ListarUsuarios(100, 0)
	if err != nil {
		log.Printf("fallo al leer usuarios, con error: %s", err.Error())
		http.Error(writer, "fallo al leer los usuarios", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(usuarios)
}

func (h *HandlerUsuarios) CrearUsuario(writer http.ResponseWriter, req *http.Request) { // http.HandlerFunc {
	/*
		Crea los usuarios nuevos a través de POST.
	*/
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al crear un nuevo usuario, con error: %s", err.Error())
		http.Error(writer, "fallo al crear un nuevo usuario", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	nuevoId, err := h.controller.CrearUsuario(body)
	if err != nil {
		log.Println("fallo al crear un nuevo usuario, con error:", err.Error())
		http.Error(writer, "fallo al crear un nuevo usuario", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(nuevoId))
}

func (h *HandlerUsuarios) ActualizarUsuario(writer http.ResponseWriter, req *http.Request) {
	/*
		Función que permite actualizar la información asociada a un usuario existente en
		la base de datos de usuarios.
	*/
	vars := mux.Vars(req)
	usu := vars["usu"]
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un usuario, con error: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	err = h.controller.ActualizarUnUsuario(body, usu)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		http.Error(writer, fmt.Sprintf("fallo al actualizar un usuario, con error: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (h *HandlerUsuarios) TraerUsuario(writer http.ResponseWriter, req *http.Request) {
	/*
		Valida los usuarios por el usu asociado.
	*/
	vars := mux.Vars(req)
	usu := vars["usu"]

	usuario, err := h.controller.LeerUnUsuario(usu)
	if err != nil {
		log.Printf("fallo al leer un usuario, con error: %s", err.Error())
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(fmt.Sprintf("el usuario con usu %s no se pudo encontrar", usu)))
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(usuario)
}
