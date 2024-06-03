package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	models "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/models"
	repository "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/repository"
)

/*
	Funciones encargadas de implementar las funciones relacionadas a la base de datos de los usuarios
*/

//var hc *HandlerUsuarios

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

func (hc *HandlerUsuarios) ActualizarUsuario() http.HandlerFunc {
	/*
		Función de registro, permite crear nuevos usuarios en la base de datos de usuarios.
		Usu         (string, nombre del usuario)
		Password    (string, contraseña del usuario)
		Automobiles (int, número de automobiles que el usuario ha rentado)
		Types_cars  ([]string, slice con la referencia de cada uno de los carros prestados)
		Debts (int, deuda del usuario)
	*/
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
		} /*else if change == "reservation" {
			//MEJOR HACER ESTO POR HILOS (QUE SE EJECUTE EN AMBAS PARTES AL TIEMPO AUTOS Y USERS)
			var usuBase models.User = hc.BD.Memoria[usuario.Usu] //se altera en base usuario

			if usuBase.Automobiles != 0 { //antes se tenían rentados otros carros
				//se borran todos los carros que se tenían:
				for _, ref_car := range usuBase.Types_cars {
					var carBase models.Automobile = hc_car.BD.Memoria[ref_car]
					carBase.Usedby = ""
					carBase.DateBegin = time.Date(0001, 1, 1, 00, 00, 00, 00, time.UTC)
					carBase.DateEnd = time.Date(0001, 1, 1, 00, 00, 00, 00, time.UTC)
					hc_car.BD.Memoria[ref_car] = carBase
				}
				//
			}
			hc.BD.Memoria[usuario.Usu] = usuBase
		}*/
		w.WriteHeader(http.StatusCreated)
	})
}

func (hc *HandlerUsuarios) TraerUsuario() http.HandlerFunc {
	/*
		Función de validación, permite retornar información especifica (toda) de un usuario.
	*/
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
