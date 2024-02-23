package main

import (
	"github.com/DiegoADG/challengemelidiegoduran/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/obtenerArchivo", controller.ObtenerArchivo)
	router.GET("/consultaritems", controller.ConsultarItems)

	router.Run()

}
