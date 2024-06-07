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

	/* router (multiplexador) a los endpoints de la API (implementado con el paquete gorilla/mux) */
	router := mux.NewRouter()

	/* rutas a los endpoints de la API */
	router.Handle("/users", http.HandlerFunc(handler_user.ListarUsuarios)).Methods(http.MethodGet)
	router.Handle("/users", http.HandlerFunc(handler_user.CrearUsuario)).Methods(http.MethodPost)
	//router.Handle("/posts/{id}", http.HandlerFunc(handler_user.TraerComentario)).Methods(http.MethodGet)
	//router.Handle("/posts/{id}", http.HandlerFunc(handler_user.ActualizarComentario)).Methods(http.MethodPatch)
	//router.Handle("/posts/{id}", http.HandlerFunc(handler_user.EliminarComentario)).Methods(http.MethodDelete)

	/* servidor escuchando en localhost por el puerto 8080 y entrutando las peticiones con el router */
	log.Fatal(http.ListenAndServe(":8080", router))
}

/*
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
*/
