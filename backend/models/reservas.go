package models

type Reserva struct {
	IdUser  string              `json: "iduser"` //nombre del usuario (Usu      string `json: "usu"`)
	Details map[string]struct { //por cada referencia seleccionada
		//Ref            string `json: "ref"`   //tipo de referencia (key del mapa)
		Total          int  `json: "total"` //se calcula con Automobile.Price * Days
		Days           int  `json: "days"`  //dias de renta
		LifeInsurance  bool `json: "lifeinsurance"`
		RoadAssistance bool `json: "roadassistance"`
		BabySeat       bool `json: "babyseat"`
	} `json: "details"`
}

/*
Notas: la eliminacion de reserva es por referencia seleccionada
*/
