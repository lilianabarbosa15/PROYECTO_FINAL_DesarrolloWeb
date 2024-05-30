package main

import (
	"log"
	"net/http"

	handlers "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/handlers"
	repository "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/repository"
)

func main() {
	bd := repository.NewBaseDatosUsers()
	handler := handlers.NewHandlerUsuarios(bd)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", handler.ListarUsuarios())
	mux.HandleFunc("POST /users/new", handler.NuevoUsuario())
	mux.HandleFunc("GET /users/{usu}", handler.TraerUsuario())
	//mux.HandleFunc("GET /users/report/{usu}", handler.ReportedeUsuario())

	//Escucha del puerto local :8080
	log.Fatal(http.ListenAndServe(":8080", mux))
}
