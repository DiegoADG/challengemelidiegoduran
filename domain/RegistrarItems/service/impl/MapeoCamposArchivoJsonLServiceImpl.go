package service

import (
	"encoding/json"
	"fmt"
)

type MapeoCamposArchivoServiceJsonLImpl struct {
	Site string `json:"site"`
	Id   string `json:"id"`
}

func (M MapeoCamposArchivoServiceJsonLImpl) MapeoCamposArchivo(lineaArchivo string, separador string) []string {
	var lineaJson MapeoCamposArchivoServiceJsonLImpl
	err := json.Unmarshal([]byte(lineaArchivo), &lineaJson)
	fmt.Printf("ID: " + lineaJson.Id)
	fmt.Printf("Site: " + lineaJson.Site)
	if err != nil {
		panic(err)
	}
	var dato []string = []string{lineaJson.Site, lineaJson.Id}
	return dato

}
