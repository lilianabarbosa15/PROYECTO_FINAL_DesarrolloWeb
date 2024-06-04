package models

type Automobile struct {
	Ref               string `json: "ref"`               //referencia del carro en stock
	Type_transmission string `json: "type_transmission"` //puede ser: manual ó automático
	Type_fuel         string `json: "type_fuel"`         //puede ser: gasolina, Diesel ó eléctrico
	Year              int    `json: "year"`              //año de fabricación del modelo del carro
	Model             string `json: "model"`             //modelo del carro en stock
	Color             string `json: "color"`             //color del carro en stock
	Price             int    `json: "price"`             //costo de renta del carro por día
	Seats             int    `json: "seats"`             //capacidad de personas que soporta el carro
	Brand             string `json: "brand"`             //marca del carro en stock
	Image             string `json: "image"`             //link de la imagen del carro
	Quantity          int    `json: "quantity"`          // diponibilidad (stock) (lo unico que va alterando es eso)
}
