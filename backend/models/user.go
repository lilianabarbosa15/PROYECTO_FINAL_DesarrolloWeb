package models

type User struct {
	Usu      string `json: "usu"`      //nombre del usuario
	Name     string `json: "name"`     //nombre de la persona
	Email    string `json: "email"`    //contraseña del usuario
	Password string `json: "password"` //contraseña del usuario
}
