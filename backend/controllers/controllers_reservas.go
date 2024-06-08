package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/models"
	repositorio "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/repository"
)

var (
	updateReservaQuery = "UPDATE reservas SET %s WHERE id=:id;"
	deleteReservaQuery = "DELETE FROM reservas WHERE id=$1;"
	listReservaQuery   = "SELECT * FROM reservas limit $1 offset $2"
	createReservaQuery = "INSERT INTO reservas (iduser, ref, total, days, lifeinsurance, roadassistance, babyseat, luxury) VALUES (:iduser, :ref, :total, :days, :lifeinsurance, :roadassistance, :babyseat, :luxury) returning id;"
)

type ReservaController struct {
	repo repositorio.Repository[models.Reserva]
}

func NewReservaController(repo repositorio.Repository[models.Reserva]) (*ReservaController, error) {
	if repo == nil {
		return nil, fmt.Errorf("para instanciar un controlador se necesita un repositorio no nulo")
	}
	return &ReservaController{
		repo: repo,
	}, nil
}

func (c *ReservaController) ListarReservas(limit, offset int) ([]byte, error) {
	reservas, _, err := c.repo.List(context.TODO(), listReservaQuery, limit, offset)
	fmt.Println("reservas: ", reservas)
	if err != nil {
		log.Printf("fallo al leer reservas, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer reservas, con error: %s", err.Error())
	}
	for _, detalle := range reservas {
		fmt.Println("IdUser: ", detalle.IdUser)
		fmt.Println("Ref: ", detalle.Ref)
		fmt.Println("Total: ", detalle.Total)
		fmt.Println("Days: ", detalle.Days)
		fmt.Println("LifeInsurance: ", detalle.LifeInsurance)
		fmt.Println("RoadAssistance: ", detalle.RoadAssistance)
		fmt.Println("BabySeat: ", detalle.BabySeat)
	}
	jsonReservas, err := json.Marshal(reservas)
	if err != nil {
		log.Printf("fallo al leer reservas, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer reservas, con error: %s", err.Error())
	}
	return jsonReservas, nil
}

func (c *ReservaController) CrearReserva(reqBody []byte) (int64, error) {
	nuevaReserva := &models.Reserva{}
	err := json.Unmarshal(reqBody, nuevaReserva)
	if err != nil {
		log.Printf("fallo al crear nueva reserva, con error: %s", err.Error())
		return 0, fmt.Errorf("fallo al crear nueva reserva, con error: %s", err.Error())
	}

	valoresColumnasNuevaReserva := map[string]any{
		"iduser":         nuevaReserva.IdUser,
		"ref":            nuevaReserva.Ref,
		"total":          nuevaReserva.Total,
		"days":           nuevaReserva.Days,
		"lifeinsurance":  nuevaReserva.LifeInsurance,
		"roadassistance": nuevaReserva.RoadAssistance,
		"babyseat":       nuevaReserva.BabySeat,
	}
	id, er := c.repo.Create_id(context.TODO(), createReservaQuery, valoresColumnasNuevaReserva)
	if er != nil {
		log.Printf("fallo al crear nueva reserva, con error: %s", er.Error())
		return 0, fmt.Errorf("fallo al crear nueva reserva, con error: %s", er.Error())
	}
	return id, nil
}

func (c *ReservaController) BorrarUnaReserva(reqBody []byte, usu, ref string) error {

	reservas, _, err := c.repo.List(context.TODO(), listReservaQuery, 10000, 0)
	fmt.Println("reservas: ", reservas)
	if err != nil {
		log.Printf("fallo al leer reservas, con error: %s", err.Error())
		return fmt.Errorf("fallo al leer reservas, con error: %s", err.Error())
	}
	for _, detalle := range reservas {
		if detalle.IdUser == usu && detalle.Ref == ref {
			err := c.repo.Delete(context.TODO(), deleteReservaQuery, strconv.Itoa(detalle.Id))
			if err != nil {
				log.Printf("fallo al eliminar un comentario, con error: %s", err.Error())
				return fmt.Errorf("fallo al eliminar un comentario, con error: %s", err.Error())
			}
			return nil
		}
	}
	return nil
}

func (c *ReservaController) ActualizarUnaReserva(reqBody []byte, usu, ref string) error {
	reservas, _, er := c.repo.List(context.TODO(), listReservaQuery, 10000, 0)
	fmt.Println("reservas: ", reservas)
	if er != nil {
		log.Printf("fallo al leer reservas, con error: %s", er.Error())
		return fmt.Errorf("fallo al leer reservas, con error: %s", er.Error())
	}
	var id int
	for _, detalle := range reservas {
		if detalle.IdUser == usu && detalle.Ref == ref {
			id = detalle.Id
		}
	}

	nuevosValoresReserva := make(map[string]any)
	err := json.Unmarshal(reqBody, &nuevosValoresReserva)
	if err != nil {
		log.Printf("fallo al actualizar una reserva, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar una reserva, con error: %s", err.Error())
	}

	if len(nuevosValoresReserva) == 0 {
		log.Printf("fallo al actualizar una reserva, con error: no hay datos")
		return fmt.Errorf("fallo al actualizar una reserva, con error: no hay datos")
	}

	query := construirUpdateReservaQuery(nuevosValoresReserva)
	nuevosValoresReserva["id"] = id
	err = c.repo.Update(context.TODO(), query, nuevosValoresReserva)
	if err != nil {
		log.Printf("fallo al actualizar una reserva, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar una reserva, con error: %s", err.Error())
	}
	return nil
}

func construirUpdateReservaQuery(nuevosValores map[string]any) string {
	columnas := []string{}
	for key := range nuevosValores {
		columnas = append(columnas, fmt.Sprintf("%s=:%s", key, key))
	}
	columnasString := strings.Join(columnas, ",")
	return fmt.Sprintf(updateReservaQuery, columnasString)
}
