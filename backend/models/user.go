package models

type User struct {
	Usu      string `json:"usu" db:"usu"`           //nombre del usuario
	Name     string `json:"name" db:"name"`         //nombre de la persona
	Email    string `json:"email" db:"email"`       //contraseña del usuario
	Password string `json:"password" db:"password"` //contraseña del usuario
}
