package models

type Automobile struct {
	Ref               string `json: "ref" db: "ref"`                             //referencia del carro en stock
	Type_transmission string `json: "type_transmission" db: "type_transmission"` //puede ser: manual ó automático
	Type_fuel         string `json: "type_fuel" db: "type_fuel"`                 //puede ser: gasolina, Diesel ó eléctrico
	Year              int    `json: "year" db: "year"`                           //año de fabricación del modelo del carro
	Model             string `json: "model" db: "model"`                         //modelo del carro en stock
	Color             string `json: "color" db: "color"`                         //color del carro en stock
	Price             int    `json: "price" db: "price"`                         //costo de renta del carro por día
	Seats             int    `json: "seats" db: "seats"`                         //capacidad de personas que soporta el carro
	Brand             string `json: "brand" db: "brand"`                         //marca del carro en stock
	Image             string `json: "image" db: "image"`                         //link de la imagen del carro
	Quantity          int    `json: "quantity" db: "quantity"`                   // diponibilidad (stock) (lo unico que va alterando es eso)
}
