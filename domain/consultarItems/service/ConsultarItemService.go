package service

import (
	dto "github.com/DiegoADG/challengemelidiegoduran/domain/consultarItems/dto"
)

type ConsultarItems interface {
	ConsultarItems() dto.RespuestadeItemsdto
}
