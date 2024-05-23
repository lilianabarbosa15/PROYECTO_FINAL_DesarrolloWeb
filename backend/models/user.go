package models

type User struct {
	Usu         string   //nombre del usuario
	Password    string   //contraseña del usuario
	Automobiles int      //número de automobiles que el usuario ha rentado
	Types_cars  []string //slice con la referencia de cada uno de los carros prestados
	Debts       int      //deuda del usuario (dollars)
}
