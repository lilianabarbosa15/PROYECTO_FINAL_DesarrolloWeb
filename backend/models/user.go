package models

import "time"

type User struct {
	Usu         string                 `json: "usu"`         //nombre del usuario
	Password    string                 `json: "password"`    //contraseña del usuario
	Automobiles int                    `json: "automobiles"` //número de automobiles que el usuario ha rentado
	Types_cars  map[string][]time.Time `json: "types_cars"`  //slice con la referencia de cada uno de los carros prestados
	Debts       int                    `json: "debts"`       //deuda del usuario (dollars)
}
