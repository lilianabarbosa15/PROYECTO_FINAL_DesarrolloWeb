package models

type Reserva struct {
	Id             int    `json:"id" db:"id"`
	IdUser         string `json:"iduser" db:"iduser"` //id del usuario (usu) no se repite
	Ref            string `json:"ref" db:"ref"`       //tipo de referencia (key del mapa)
	Total          int    `json:"total" db:"total"`   //se calcula con Automobile.Price * Days
	Days           int    `json:"days" db:"days"`     //dias de renta
	LifeInsurance  bool   `json:"lifeinsurance" db:"lifeinsurance"`
	RoadAssistance bool   `json:"roadassistance" db:"roadassistance"`
	BabySeat       bool   `json:"babyseat" db:"babyseat"`
	Luxury         bool   `json:"luxury" db:"luxury"`
}

/*
Notas: la eliminacion de reserva es por referencia seleccionada
*/
