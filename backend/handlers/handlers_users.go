package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	models "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/models"
	repository "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/repository"
)

/*
	Funciones encargadas de implementar las funciones relacionadas a la base de datos de los usuarios
*/

type HandlerUsuarios struct {
	BD *repository.BaseDatosUsers
}

func NewHandlerUsuarios(bd *repository.BaseDatosUsers) *HandlerUsuarios {
	return &HandlerUsuarios{
		BD: bd,
	}
}

func (hc *HandlerUsuarios) ListarUsuarios() http.HandlerFunc {
	/*
		Función que retorna toda la información de la base de datos de usuarios.
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users := []models.User{}
		for _, user := range hc.BD.Memoria {
			users = append(users, user)
		}
		jsonUsu, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "fallo al comunicar en json", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonUsu)
	})
}

func (hc *HandlerUsuarios) NuevoUsuario() http.HandlerFunc {
	/*
		Función de registro, permite crear nuevos usuarios en la base de datos de usuarios.
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "fallo la peticion POST", http.StatusBadRequest)
		}
		usuario := models.User{}
		err = json.Unmarshal(body, &usuario)
		if err != nil {
			http.Error(w, "fallo al codificar en json", http.StatusInternalServerError)
		}
		hc.BD.Memoria[usuario.Usu] = usuario
		w.WriteHeader(http.StatusCreated)
	})
}

func (hc *HandlerUsuarios) TraerUsuario() http.HandlerFunc {
	/*
		Función de validación, permite retornar información especifica de un usuario.
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usu := r.PathValue("usu")
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

/*func (hc *HandlerUsuarios) ReportedeUsuario() http.HandlerFunc {
	/*
		Función de .
	//
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//crea la reserva en la base de datos de los carros:

		//retorna la lista de autos a nombre del usuario:
		usu := r.PathValue("usu")
		if usu == "" {
			http.Error(w, "usuario no valido", http.StatusBadRequest)
			return
		}
		infousuario, ok := hc.BD.Memoria[usu]
		if !ok {
			http.Error(w, "no se encuentra informacion para ese usuario", http.StatusNotFound)
			return
		}
		report := []string{}
		report = append(report, infousuario.Types_cars...)
		report = append(report, strconv.Itoa(infousuario.Debts))
		//
		payload, err := json.Marshal(report)
		if err != nil {
			http.Error(w, "fallo la codificacion a JSON de la informacion", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	})
}*/
