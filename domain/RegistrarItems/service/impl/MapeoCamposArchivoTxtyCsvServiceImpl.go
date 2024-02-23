package service

import (
	"strings"
)

type MapeoCamposArchivoTxtyCsvServiceImpl struct {
}

func (M MapeoCamposArchivoTxtyCsvServiceImpl) MapeoCamposArchivo(lineaArchivo string, separador string) []string {

	registroArchivo := strings.Split(lineaArchivo, separador)

	return registroArchivo

}
