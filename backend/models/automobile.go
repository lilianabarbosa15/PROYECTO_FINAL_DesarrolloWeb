package models

import "time"

type Automobile struct {
	Ref               string    //referencia del carro en stock
	Type_transmission string    //puede ser: manual ó automático
	Type_fuel         string    //puede ser: gasolina, Diesel ó eléctrico
	Model             string    //modelo del carro en stock
	Color             string    //color del carro en stock
	Price             int       //costo de renta del carro por día
	Seats             int       //capacidad de personas que soporta el carro
	Brand             string    //marca del carro en stock
	Image             string    //link de la imagen del carro
	Daysinuse         int       //días que va a estar prestado el carro sino está en prestado se pone 0
	Usedby            string    //nombre de la persona que presta sino está en prestamo se pone ""
	Available         time.Time //fecha de retorno del carro si está prestado
}
