package consultarcategoria

import (
	"github.com/DiegoADG/challengemelidiegoduran/integration/consultarCategoria/dto"
)

type ConsultarCategoriaService interface {
	ConsultarCategorias(authorization string, idCategoria string) dto.ConsultarCategoriaDto
}
