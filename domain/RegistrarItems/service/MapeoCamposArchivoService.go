package domain

type MapeoCamposArchivoService interface {
	MapeoCamposArchivo(lineaArchivo string, separador string) []string
}
