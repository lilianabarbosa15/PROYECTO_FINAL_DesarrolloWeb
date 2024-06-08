package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/models"
	repositorio "github.com/lilianabarbosa15/PROYECTO_FINAL_DesarrolloWeb/repository"
)

var (
	updateUserQuery = "UPDATE users SET %s WHERE usu=:usu;"
	selectUserQuery = "SELECT usu, name, email, password FROM users WHERE usu=$1;"
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
		log.Printf("fallo al crear un nuevo usuario, con error: %s", err.Error())
		return "", fmt.Errorf("fallo al crear un nuevo usuario, con error: %s", err.Error())
	}

	valoresColumnasNuevoUsuario := map[string]any{
		"usu":      nuevoUsuario.Usu,
		"name":     nuevoUsuario.Name,
		"email":    nuevoUsuario.Email,
		"password": nuevoUsuario.Password,
	}
	er := c.repo.Create(context.TODO(), createUserQuery, valoresColumnasNuevoUsuario)
	if er != nil {
		log.Printf("fallo al crear un nuevo usuario, con error: %s", er.Error())
		return nuevoUsuario.Usu, fmt.Errorf("fallo al crear un nuevo usuario, con error: %s", er.Error())
	}
	return nuevoUsuario.Usu, nil
}

func (c *UserController) ActualizarUnUsuario(reqBody []byte, usu string) error {
	//PATCH
	nuevosValoresUsuario := make(map[string]any)
	err := json.Unmarshal(reqBody, &nuevosValoresUsuario)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un usuario, con error: %s", err.Error())
	}

	if len(nuevosValoresUsuario) == 0 {
		log.Printf("fallo al actualizar un usuario, con error: no hay datos")
		return fmt.Errorf("fallo al actualizar un usuario, con error: no hay datos")
	}
	query := construirUpdateQuery(nuevosValoresUsuario)
	nuevosValoresUsuario["usu"] = usu
	err = c.repo.Update(context.TODO(), query, nuevosValoresUsuario)
	if err != nil {
		log.Printf("fallo al actualizar un usuario, con error: %s", err.Error())
		return fmt.Errorf("fallo al actualizar un usuario, con error: %s", err.Error())
	}
	return nil
}

func construirUpdateQuery(nuevosValores map[string]any) string {
	columnas := []string{}
	for key := range nuevosValores {
		columnas = append(columnas, fmt.Sprintf("%s=:%s", key, key))
	}
	columnasString := strings.Join(columnas, ",")
	return fmt.Sprintf(updateUserQuery, columnasString)
}

func (c *UserController) LeerUnUsuario(usu string) ([]byte, error) {
	usuario, err := c.repo.Read(context.TODO(), selectUserQuery, usu)
	if err != nil {
		log.Printf("fallo al leer un usuario, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un usuario, con error: %s", err.Error())
	}
	usuarioJson, err := json.Marshal(usuario)
	if err != nil {
		log.Printf("fallo al leer un usuario, con error: %s", err.Error())
		return nil, fmt.Errorf("fallo al leer un usuario, con error: %s", err.Error())
	}
	return usuarioJson, nil
}
