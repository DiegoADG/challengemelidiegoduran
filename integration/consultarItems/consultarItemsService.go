package consultarItems

import (
	"github.com/DiegoADG/challengemelidiegoduran/integration/consultarItems/dto"
)

type ConsultarItemsService interface {
	ConsultarItems(authorization string, idItem string) dto.ConsultarItemsDto
}
