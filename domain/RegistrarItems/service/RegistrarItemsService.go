package domain

import (
	"bufio"

	"github.com/DiegoADG/challengemelidiegoduran/domain/RegistrarItems/dto"
)

type ObtenerArchivo interface {
	RegistrarItems(archivo *bufio.Scanner, autorizacion string, separador string, extension string) dto.ObtenerArchivoResponseDto
}
