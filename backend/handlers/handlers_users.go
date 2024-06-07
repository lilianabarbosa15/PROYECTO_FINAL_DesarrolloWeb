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
	controller *controllers.UserController // *repository.BaseDatosUsers
}

func NewHandlerUsuarios(controller *controllers.UserController) (*HandlerUsuarios, error) { // *repository.BaseDatosUsers) *HandlerUsuarios {
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
		.
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
	writer.Write([]byte(nuevoId)) //[]byte(fmt.Sprintf("id nuevo comentario: %d", nuevoId)))
}

func (h *HandlerUsuarios) ActualizarUsuario(writer http.ResponseWriter, req *http.Request) {
	/*
		Función que permite actualizar la información
		asociada a un usuario existente en la base de datos de usuarios.
		Usu         (string, nombre del usuario)
		Password    (string, contraseña del usuario)
		Automobiles (int, número de automobiles que el usuario ha rentado)
		Types_cars  ([]string, slice con la referencia de cada uno de los carros prestados)
		Debts (int, deuda del usuario)
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

/*
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		change := mux.Vars(r)["change"]
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "fallo la peticion POST", http.StatusBadRequest)
		}
		usuario := models.User{}
		err = json.Unmarshal(body, &usuario)
		if err != nil {
			http.Error(w, "fallo al codificar en json", http.StatusInternalServerError)
		}
		if change == "registration" {
			hc.BD.Memoria[usuario.Usu] = usuario
		} else if change == "information" {
			var usuBase models.User = hc.BD.Memoria[usuario.Usu]
			usuBase.Password = usuario.Password
			hc.BD.Memoria[usuario.Usu] = usuBase
		}
		w.WriteHeader(http.StatusCreated)
	})
}*/

/*
func (hc *HandlerUsuarios) TraerUsuario() http.HandlerFunc {
	/*
		Función de validación, permite retornar información especifica (toda) de un usuario.
	//
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usu := mux.Vars(r)["usu"]
		if usu == "" {
			http.Error(w, "usuario no valido", http.StatusBadRequest)
			return
		}
		infousuario, ok := hc.BD.Memoria[usu]
		if !ok {
			http.Error(w, "no se encuentra informacion para ese usuario", http.StatusNotFound)
			return
		}
		payload, err := json.Marshal(infousuario)
		if err != nil {
			http.Error(w, "fallo la codificacion a JSON de la informacion", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	})
}
*/
