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
	bd_reserva := repository.NewBaseDatosReservas()
	handler_reservas := handlers.NewHandlerReservas(bd_reserva)

	// Se agrega multiplexor y enrutador:
	mux := mux.NewRouter()
	//asociar handlers al mux:
	//usuarios
	mux.HandleFunc("/users", handler_usu.ListarUsuarios()).Methods("GET")                        /////////////////////////////////////////
	mux.HandleFunc("/users/{change}/", handler_usu.ActualizarUsuario()).Methods("POST", "PATCH") //crear, cambiar contraseña y/o reserva
	mux.HandleFunc("/validation/{usu}/", handler_usu.TraerUsuario()).Methods("GET")              //trae la información de un usuario en especifico
	//autos
	mux.HandleFunc("/automobiles", handler_car.ListarAutos()).Methods("GET")                       //permite traer toda la información asociada a la base de datos de vehiculos
	mux.HandleFunc("/automobiles", handler_car.NuevoAuto()).Methods("POST")                        //crea nuevo auto en la base de datos de automobiles
	mux.HandleFunc("/automobiles/{ref}/{quantity}", handler_car.ActualizarAuto()).Methods("PATCH") //crea nuevo auto en la base de datos de automobiles
	mux.HandleFunc("/automobiles/{filter}/{kind}", handler_car.TraerAutos()).Methods("GET")        //filtrado de la base de datos de los autos
	//reservas
	mux.HandleFunc("/reservations", handler_reservas.ListarReservas()).Methods("GET")               //lista todas las reservas que se tienen en la base de datos
	mux.HandleFunc("/reservations", handler_reservas.NuevaReserva()).Methods("POST")                //crea una nueva reserva desde cero (agregando un nuevo usuario)
	mux.HandleFunc("/reservations", handler_reservas.BorrarReserva()).Methods("DELETE")             //elimina totalmente una reserva de la base de datos
	mux.HandleFunc("/reservations/{change}", handler_reservas.ActualizarReserva()).Methods("PATCH") //edita parcialmente los detalles de la reserva de deteminado usuario

	// Definición de servidor que esté escuchando:
	log.Fatal(http.ListenAndServe(":8080", mux))
}
