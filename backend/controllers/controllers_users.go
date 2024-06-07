package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	_ "strings"

	"github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/models"
	repositorio "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/repository"
)

var (
	//updateQuery     = "UPDATE comentarios SET %s WHERE id=:id;"
	//deleteQuery     = "DELETE FROM comentarios WHERE id=$1;"
	//selectQuery     = "SELECT id, time, comment, reactions FROM comentarios WHERE id=$1;"
	listUserQuery   = "SELECT usu, name, email, password FROM users limit $1 offset $2"
	createUserQuery = "INSERT INTO users (usu, name, email, password) VALUES (:usu, :name, :email, :password);" // return usu;" ////////////////////
)

type UserController struct {
	repo repositorio.Repository[models.User]
}

func NewUserController(repo repositorio.Repository[models.User]) (*UserController, error) {
	if repo == nil {
		return nil, fmt.Errorf("para instanciar un controlador se necesita un repositorio no nulo")
	}
	return &UserController{
		repo: repo,
	}, nil
}

func (c *UserController) ListarUsuarios(limit, offset int) ([]byte, error) {
	fmt.Println("limit: ", limit)
	fmt.Println("offset: ", offset)
	fmt.Println("context.TODO(): ", context.TODO())
	usuarios, _, err := c.repo.List(context.TODO(), listUserQuery, limit, offset)
	if err != nil {
		log.Printf("fallo al leer usuarios, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer usuarios, con error: %s", err.Error())
	}

	jsonUsuarios, err := json.Marshal(usuarios)
	if err != nil {
		log.Printf("fallo al leer usuarios, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer usuarios, con error: %s", err.Error())
	}
	return jsonUsuarios, nil
}

func (c *UserController) CrearUsuario(reqBody []byte) (string, error) {
	nuevoUsuario := &models.User{}
	err := json.Unmarshal(reqBody, nuevoUsuario)
	if err != nil {
		log.Printf("fallo al crear un nuevo comentario, con error: %s", err.Error())
		return "", fmt.Errorf("fallo al crear un nuevo comentario, con error: %s", err.Error())
	}

	valoresColumnasNuevoUsuario := map[string]any{
		"usu":      nuevoUsuario.Usu,
		"name":     nuevoUsuario.Name,
		"email":    nuevoUsuario.Email,
		"password": nuevoUsuario.Password,
	}

	//DEVUELVE ID NUMERICO///////////////////////////////////////////////////////////////////////////////////////////
	er := c.repo.Create(context.TODO(), createUserQuery, valoresColumnasNuevoUsuario)
	if er != nil {
		log.Printf("fallo al crear un nuevo comentario, con error: %s", er.Error())
		return nuevoUsuario.Usu, fmt.Errorf("fallo al crear un nuevo comentario, con error: %s", er.Error())
	}
	return nuevoUsuario.Usu, nil
}

/*func (c *Controller) ActualizarUnComentario(reqBody []byte, id string) error {
	//PATCH
	nuevosValoresComentario := make(map[string]any)
	err := json.Unmarshal(reqBody, &nuevosValoresComentario)
	if err != nil {
		log.Printf("fallo al actualizar un comentario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un comentario, con error: %s", err.Error())
	}

	if len(nuevosValoresComentario) == 0 {
		log.Printf("fallo al actualizar un comentario, con error: no hay datos")
		return fmt.Errorf("fallo al actualizar un comentario, con error: no hay datos")
	}

	query := construirUpdateQuery(nuevosValoresComentario)
	nuevosValoresComentario["id"] = id
	err = c.repo.Update(context.TODO(), query, nuevosValoresComentario)
	if err != nil {
		log.Printf("fallo al actualizar un comentario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un comentario, con error: %s", err.Error())
	}
	return nil
}*/

/*func construirUpdateQuery(nuevosValores map[string]any) string {
	columnas := []string{}
	for key := range nuevosValores {
		columnas = append(columnas, fmt.Sprintf("%s=:%s", key, key))
	}
	columnasString := strings.Join(columnas, ",")
	return fmt.Sprintf(updateQuery, columnasString)
}*/
