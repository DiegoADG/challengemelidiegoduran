package controller

import (
	consultarDto "github.com/DiegoADG/challengemelidiegoduran/domain/consultarItems/dto"
	consultar "github.com/DiegoADG/challengemelidiegoduran/domain/consultarItems/service"
	consultarImpl "github.com/DiegoADG/challengemelidiegoduran/domain/consultarItems/service/impl"
	"github.com/gin-gonic/gin"
)

func consultaItems(consultar consultar.ConsultarItems) consultarDto.RespuestadeItemsdto {
	return consultar.ConsultarItems()
}

func ConsultarItems(c *gin.Context) {
	consultar := consultarImpl.ConsultarItemsServiceImpl{}
	items := consultaItems(consultar)

	c.JSON(200, items)
}
