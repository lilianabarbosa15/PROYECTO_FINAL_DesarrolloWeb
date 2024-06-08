package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/controllers"
	"github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/handlers"
	"github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/models"
	repositorio "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/repository" /* importando el paquete de repositorio */
)

/*
función para conectarse a la instancia de PostgreSQL, en general sirve para cualquier base de datos SQL.
Necesita la URL del host donde está instalada la base de datos y el tipo de base datos (driver)
*/
func ConectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl) // driver: postgres
	if err != nil {
		log.Printf("fallo la conexion a PostgreSQL, error: %s", err.Error())
		return nil, err
	}

	log.Printf("Nos conectamos bien a la base de datos db: %#v", db)
	return db, nil
}

func main() {
	/* creando un objeto de conexión a PostgreSQL */
	db, err := ConectarDB(fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")), "postgres")
	fmt.Printf("postgres://%s:%s@db:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
		return
	}

	//usuarios:
	repo_user, err := repositorio.NewRepository[models.User](db)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
		return
	}
	controller_user, err := controllers.NewUserController(repo_user)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de controller", err.Error())
		return
	}
	handler_user, err := handlers.NewHandlerUsuarios(controller_user)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
		return
	}
	//automobiles:
	repo_auto, err := repositorio.NewRepository[models.Automobile](db)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
		return
	}
	controller_auto, err := controllers.NewAutoController(repo_auto)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de controller", err.Error())
		return
	}
	handler_auto, err := handlers.NewHandlerAutomobiles(controller_auto)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
		return
	}
	//reservas:
	repo_reserva, err := repositorio.NewRepository[models.Reserva](db)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de repositorio", err.Error())
		return
	}
	controller_reserva, err := controllers.NewReservaController(repo_reserva)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de controller", err.Error())
		return
	}
	handler_reservas, err := handlers.NewHandlerReservas(controller_reserva)
	if err != nil {
		log.Fatalln("fallo al crear una instancia de handler", err.Error())
		return
	}

	/* router (multiplexador) a los endpoints de la API (implementado con el paquete gorilla/mux) */
	router := mux.NewRouter()

	/* rutas a los endpoints de la API */
	//usuarios:
	router.Handle("/users", http.HandlerFunc(handler_user.ListarUsuarios)).Methods(http.MethodGet)
	router.Handle("/users", http.HandlerFunc(handler_user.CrearUsuario)).Methods(http.MethodPost)             //crea nuevo usuario
	router.Handle("/users/{usu}", http.HandlerFunc(handler_user.ActualizarUsuario)).Methods(http.MethodPatch) //cambiar contraseña y/o reserva
	router.Handle("/validation/{usu}", http.HandlerFunc(handler_user.TraerUsuario)).Methods(http.MethodGet)   //trae la información de un usuario en especifico

	//autos:
	router.Handle("/automobiles", http.HandlerFunc(handler_auto.ListarAutos)).Methods(http.MethodGet)                //permite traer toda la información asociada a la base de datos de vehiculos
	router.Handle("/automobiles", http.HandlerFunc(handler_auto.NuevoAuto)).Methods(http.MethodPost)                 //crea nuevo auto en la base de datos de automobiles
	router.Handle("/automobiles/{filter}/{kind}", http.HandlerFunc(handler_auto.TraerAutos)).Methods(http.MethodGet) //filtrado de la base de datos de los autos
	router.Handle("/automobiles/{ref}", http.HandlerFunc(handler_auto.ActualizarAuto)).Methods(http.MethodPatch)     //modifica características de cierto auto

	//reservas:
	router.Handle("/reservations", http.HandlerFunc(handler_reservas.ListarReservas)).Methods(http.MethodGet)                  //lista todas las reservas que se tienen en la base de datos
	router.Handle("/reservations", http.HandlerFunc(handler_reservas.NuevaReserva)).Methods(http.MethodPost)                   //crea una nueva reserva desde cero (agregando un nuevo usuario)
	router.Handle("/reservations/{usu}/{ref}", http.HandlerFunc(handler_reservas.BorrarReserva)).Methods(http.MethodDelete)    //elimina totalmente una reserva de la base de datos
	router.Handle("/reservations/{usu}/{ref}", http.HandlerFunc(handler_reservas.ActualizarReserva)).Methods(http.MethodPatch) //edita parcialmente los detalles de la reserva de deteminado usuario

	/* servidor escuchando en localhost por el puerto 8080 y entrutando las peticiones con el router */
	log.Fatal(http.ListenAndServe(":8080", router))
}
