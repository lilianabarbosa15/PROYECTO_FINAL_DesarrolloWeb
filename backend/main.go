package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/handlers"
	repository "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/repository"
)

func main() {
	bd_usu := repository.NewBaseDatosUsers()
	handler_usu := handlers.NewHandlerUsuarios(bd_usu)
	bd_car := repository.NewBaseDatosAutomobiles()
	handler_car := handlers.NewHandlerAutos(bd_car)

	// Se agrega multiplexor y enrutador:
	mux := mux.NewRouter()
	//asociar handlers al mux
	mux.HandleFunc("/users", handler_usu.ListarUsuarios()).Methods("GET")                                   /////////////////////////////////////////
	mux.HandleFunc("/users/{change}/", handler_usu.ActualizarUsuario(handler_car)).Methods("POST", "PATCH") //crear, cambiar contraseña y/o reserva
	mux.HandleFunc("/validation/{usu}/", handler_usu.TraerUsuario()).Methods("GET")                         //trae la información de un usurio en especifico
	mux.HandleFunc("/automobiles", handler_car.ListarAutos()).Methods("GET")                                //////////////////////////////////////////
	mux.HandleFunc("/automobiles/{filter}/{kind}", handler_car.TraerAutos()).Methods("GET")                 //filtrado de la base de datos de los autos
	mux.HandleFunc("/automobiles", handler_car.NuevoAuto()).Methods("POST")                                 //crea nuevo auto en la base de datos de automobiles

	// Definición de servidor que esté escuchando:
	log.Fatal(http.ListenAndServe(":8080", mux))
}
