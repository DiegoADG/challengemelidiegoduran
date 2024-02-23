package controller

import (
	"bufio"
	"fmt"
	"log"
	"path"

	registrarDto "github.com/DiegoADG/challengemelidiegoduran/domain/RegistrarItems/dto"
	registrar "github.com/DiegoADG/challengemelidiegoduran/domain/RegistrarItems/service"
	registrarImpl "github.com/DiegoADG/challengemelidiegoduran/domain/RegistrarItems/service/impl"
	"github.com/gin-gonic/gin"
)

func RegistarItem(registrar registrar.ObtenerArchivo, archivoCargado *bufio.Scanner, autorizacion string, separador string, extension string) registrarDto.ObtenerArchivoResponseDto {
	return registrar.RegistrarItems(archivoCargado, autorizacion, separador, extension)
}

func ObtenerArchivo(c *gin.Context) {
	var response registrarDto.ObtenerArchivoResponseDto

	c.MultipartForm()
	archivoCargado, informacionArchivo, errorObtenerArchivo := c.Request.FormFile("archivo")
	autorizacion := c.GetHeader("Authorization")
	separador := c.GetHeader("separador")

	if errorObtenerArchivo != nil {
		log.Fatal(errorObtenerArchivo)
		return
	}
	scanner := bufio.NewScanner(archivoCargado)
	registrar := registrarImpl.RegistrarItemsServiceImpl{}
	extension := path.Ext(informacionArchivo.Filename)

	if extension == ".txt" || extension == ".csv" {

		response = RegistarItem(registrar, scanner, autorizacion, separador, extension)
	} else if extension == ".jsonl" {

		response = RegistarItem(registrar, scanner, autorizacion, separador, extension)

	} else {
		response.Status = 400
		response.Description = "La extensiones permitidas con .txt, .csv, .jsonl"
	}

	fmt.Println(response.Status)

	c.JSON(response.Status, response)
}
