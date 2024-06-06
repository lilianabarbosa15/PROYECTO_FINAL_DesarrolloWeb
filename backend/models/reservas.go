package models

type Reserva struct {
	IdUser  string              `json: "iduser" db: "iduser"` //nombre del usuario (Usu      string `json: "usu"`)
	Details map[string]struct { //por cada referencia seleccionada
		//Ref            string `json: "ref"`   //tipo de referencia (key del mapa)
		Total          int  `json: "total" db: "total"` //se calcula con Automobile.Price * Days
		Days           int  `json: "days" db: "days"`   //dias de renta
		LifeInsurance  bool `json: "lifeinsurance" db: "lifeinsurance"`
		RoadAssistance bool `json: "roadassistance" db: "roadassistance"`
		BabySeat       bool `json: "babyseat" db: "babyseat"`
	} `json: "details" db: "details"`
}

/*
Notas: la eliminacion de reserva es por referencia seleccionada
*/
