package repository

import (
	"github.com/DiegoADG/challengemelidiegoduran/domain/consultarItems/dto"
	"github.com/DiegoADG/challengemelidiegoduran/persistence/models"
)

type ItemRepository interface {
	Create(item models.ItemsModel) error
	Consultar() (dto.RespuestadeItemsdto, error)
}
