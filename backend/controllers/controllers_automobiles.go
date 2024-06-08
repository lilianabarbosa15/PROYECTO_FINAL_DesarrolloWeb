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
	updateAutoQuery = "UPDATE automobiles SET %s WHERE ref=:ref;"
	listAutoQuery   = "SELECT * FROM automobiles limit $1 offset $2"
	createAutoQuery = "INSERT INTO automobiles (ref, type_transmission, type_fuel, year, model, color, price, seats, brand, image, quantity) VALUES (:ref, :type_transmission, :type_fuel, :year, :model, :color, :price, :seats, :brand, :image, :quantity);"
)

type AutoController struct {
	repo_auto repositorio.Repository[models.Automobile]
}

func NewAutoController(repo repositorio.Repository[models.Automobile]) (*AutoController, error) {
	if repo == nil {
		return nil, fmt.Errorf("para instanciar un controlador se necesita un repositorio no nulo")
	}
	return &AutoController{
		repo_auto: repo,
	}, nil
}

func (c *AutoController) ListarAutos(limit, offset int) ([]byte, error) {
	fmt.Println("limit: ", limit)
	fmt.Println("offset: ", offset)
	fmt.Println("context.TODO(): ", context.TODO())
	automobiles, _, err := c.repo_auto.List(context.TODO(), listAutoQuery, limit, offset)
	if err != nil {
		log.Printf("fallo al leer automobiles, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer automobiles, con error: %s", err.Error())
	}
	jsonAutomobiles, err := json.Marshal(automobiles)
	if err != nil {
		log.Printf("fallo al leer automobiles, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer automobiles, con error: %s", err.Error())
	}
	return jsonAutomobiles, nil
}

func (c *AutoController) CrearAuto(reqBody []byte) (string, error) {
	nuevoAutomobile := &models.Automobile{}
	err := json.Unmarshal(reqBody, nuevoAutomobile)
	if err != nil {
		log.Printf("fallo al crear un nuevo automobile, con error: %s", err.Error())
		return "", fmt.Errorf("fallo al crear un nuevo automobile, con error: %s", err.Error())
	}
	valoresColumnasNuevoAutomobile := map[string]any{
		"ref":               nuevoAutomobile.Ref,
		"type_transmission": nuevoAutomobile.Type_transmission,
		"type_fuel":         nuevoAutomobile.Type_fuel,
		"year":              nuevoAutomobile.Year,
		"model":             nuevoAutomobile.Model,
		"color":             nuevoAutomobile.Color,
		"price":             nuevoAutomobile.Price,
		"seats":             nuevoAutomobile.Seats,
		"brand":             nuevoAutomobile.Brand,
		"image":             nuevoAutomobile.Image,
		"quantity":          nuevoAutomobile.Quantity,
	}
	er := c.repo_auto.Create(context.TODO(), createAutoQuery, valoresColumnasNuevoAutomobile)
	if er != nil {
		log.Printf("fallo al crear un nuevo usuario, con error: %s", er.Error())
		return nuevoAutomobile.Ref, fmt.Errorf("fallo al crear un nuevo usuario, con error: %s", er.Error())
	}
	return nuevoAutomobile.Ref, nil
}

func (c *AutoController) FiltrarAutos(filter string, kind string, limit, offset int) ([]byte, error) {
	fmt.Println("limit: ", limit)
	fmt.Println("offset: ", offset)
	fmt.Println("filter: ", filter)
	fmt.Println("kind: ", kind)
	kind_int, _ := strconv.Atoi(kind)
	automobiles, _, err := c.repo_auto.List(context.TODO(), listAutoQuery, limit, offset)
	if err != nil {
		log.Printf("fallo al leer automobiles, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer automobiles, con error: %s", err.Error())
	}
	autos := []models.Automobile{}
	for _, auto := range automobiles {
		if filter == "type_transmission" {
			if auto.Type_transmission == kind {
				autos = append(autos, *auto)
			}
		} else if filter == "type_fuel" {
			if auto.Type_fuel == kind {
				autos = append(autos, *auto)
			}
		} else if filter == "year" {
			if strconv.Itoa(auto.Year) == kind {
				autos = append(autos, *auto)
			}
		} else if filter == "model" {
			if auto.Model == kind {
				autos = append(autos, *auto)
			}
		} else if filter == "color" {
			if auto.Color == kind {
				autos = append(autos, *auto)
			}
		} else if filter == "price" {
			if auto.Price <= kind_int {
				autos = append(autos, *auto)
			}
		} else if filter == "seats" {
			if auto.Seats >= kind_int {
				autos = append(autos, *auto)
			}
		} else if filter == "brand" {
			if auto.Brand == kind {
				autos = append(autos, *auto)
			}
		}
	}
	jsonCars, err := json.Marshal(autos)
	if err != nil {
		log.Printf("fallo al leer automobiles, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer automobiles, con error: %s", err.Error())
	}
	return jsonCars, nil
}

func (c *AutoController) ActualizarUnAuto(reqBody []byte, ref string) error {
	//PATCH
	nuevosValoresAuto := make(map[string]any)
	err := json.Unmarshal(reqBody, &nuevosValoresAuto)
	if err != nil {
		log.Printf("fallo al actualizar un automobile, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un automobile, con error: %s", err.Error())
	}

	if len(nuevosValoresAuto) == 0 {
		log.Printf("fallo al actualizar un automobile, con error: no hay datos")
		return fmt.Errorf("fallo al actualizar un automobile, con error: no hay datos")
	}
	query := construirUpdateCarQuery(nuevosValoresAuto)
	nuevosValoresAuto["ref"] = ref
	err = c.repo_auto.Update(context.TODO(), query, nuevosValoresAuto)
	if err != nil {
		log.Printf("fallo al actualizar un automobile, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un automobile, con error: %s", err.Error())
	}
	return nil
}

func construirUpdateCarQuery(nuevosValores map[string]any) string {
	columnas := []string{}
	for key := range nuevosValores {
		columnas = append(columnas, fmt.Sprintf("%s=:%s", key, key))
	}
	columnasString := strings.Join(columnas, ",")
	return fmt.Sprintf(updateAutoQuery, columnasString)
}
