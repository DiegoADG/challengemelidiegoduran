package service

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/DiegoADG/challengemelidiegoduran/domain/RegistrarItems/dto"
	mapeo "github.com/DiegoADG/challengemelidiegoduran/domain/RegistrarItems/service"
	consultarCategoria "github.com/DiegoADG/challengemelidiegoduran/integration/consultarCategoria"
	dtoCategoria "github.com/DiegoADG/challengemelidiegoduran/integration/consultarCategoria/dto"
	implCategoria "github.com/DiegoADG/challengemelidiegoduran/integration/consultarCategoria/impl"
	consultarCurrency "github.com/DiegoADG/challengemelidiegoduran/integration/consultarCurrency"
	dtoCurrency "github.com/DiegoADG/challengemelidiegoduran/integration/consultarCurrency/dto"
	implCurrency "github.com/DiegoADG/challengemelidiegoduran/integration/consultarCurrency/impl"
	consultarItems "github.com/DiegoADG/challengemelidiegoduran/integration/consultarItems"
	dtoItems "github.com/DiegoADG/challengemelidiegoduran/integration/consultarItems/dto"
	implItems "github.com/DiegoADG/challengemelidiegoduran/integration/consultarItems/impl"
	consultarUser "github.com/DiegoADG/challengemelidiegoduran/integration/consultarUser"
	dtoUser "github.com/DiegoADG/challengemelidiegoduran/integration/consultarUser/dto"
	implUser "github.com/DiegoADG/challengemelidiegoduran/integration/consultarUser/impl"
	modelsItems "github.com/DiegoADG/challengemelidiegoduran/persistence/models"
	repositoryItems "github.com/DiegoADG/challengemelidiegoduran/persistence/repository"
	implrepositoryItems "github.com/DiegoADG/challengemelidiegoduran/persistence/repository/impl"
)

type RegistrarItemsServiceImpl struct {
}

func consultaritem(consulta consultarItems.ConsultarItemsService, authorization string, idItem string) dtoItems.ConsultarItemsDto {
	items := consulta.ConsultarItems(authorization, idItem)
	return items
}
func consultarCategorias(consulta consultarCategoria.ConsultarCategoriaService, authorization string, idCategoria string, categoria chan dtoCategoria.ConsultarCategoriaDto) {

	categoria <- consulta.ConsultarCategorias(authorization, idCategoria)

}
func consultarCurrencys(consulta consultarCurrency.ConsultarCurrencyService, authorization string, idCurrency string, currency chan dtoCurrency.ConsultarCurrencyDto) {

	currency <- consulta.ConsultarCurrencys(authorization, idCurrency)

}
func consultarUsers(consulta consultarUser.ConsultarUserService, authorization string, idSeller int, idSite string, user chan dtoUser.ConsultarUserDto) {

	user <- consulta.ConsultarUsers(authorization, strconv.Itoa(idSeller), idSite)

}
func createItem(repository repositoryItems.ItemRepository, item modelsItems.ItemsModel) {

	err := repository.Create(item)
	if err != nil {
		fmt.Println(err)
	}
}
func Mapeo(Mapeo mapeo.MapeoCamposArchivoService, lineaArchivo string, separador string) []string {

	return Mapeo.MapeoCamposArchivo(lineaArchivo, separador)
}

func (R RegistrarItemsServiceImpl) RegistrarItems(archivoCargado *bufio.Scanner, autorizacion string, separador string, extension string) dto.ObtenerArchivoResponseDto {
	for archivoCargado.Scan() {
		lineaArchivo := archivoCargado.Text()
		var registroArchivo []string

		if extension == ".txt" || extension == ".csv" {
			mapeo := MapeoCamposArchivoTxtyCsvServiceImpl{}
			registroArchivo = Mapeo(mapeo, lineaArchivo, separador)
		} else if extension == ".jsonl" {
			mapeo := MapeoCamposArchivoServiceJsonLImpl{}
			registroArchivo = Mapeo(mapeo, lineaArchivo, "")
		}

		go registrarItemPorId(registroArchivo, autorizacion)

	}

	responseApi := dto.ObtenerArchivoResponseDto{
		Status:      200,
		Description: "Archivo procesado",
	}
	return responseApi
}

func registrarItemPorId(registroArchivo []string, autorizacion string) {

	item := implItems.ConsultarItemsServiceImpl{}
	itemConsultado := consultaritem(item, autorizacion, registroArchivo[0]+registroArchivo[1])
	if itemConsultado.Code == 200 {
		categoria := implCategoria.ConsultarCategoriaServiceImpl{}
		categoriaResultado := make(chan dtoCategoria.ConsultarCategoriaDto)
		go consultarCategorias(categoria, autorizacion, itemConsultado.Body.Category_id, categoriaResultado)

		currency := implCurrency.ConsultarCurrencyServiceImpl{}
		currencyResultado := make(chan dtoCurrency.ConsultarCurrencyDto)
		go consultarCurrencys(currency, autorizacion, itemConsultado.Body.Currency_id, currencyResultado)

		user := implUser.ConsultarUserServiceImpl{}
		userResultado := make(chan dtoUser.ConsultarUserDto)
		go consultarUsers(user, autorizacion, itemConsultado.Body.Seller_id, registroArchivo[0], userResultado)

		categoriaConsultada := <-categoriaResultado
		currencyConsultada := <-currencyResultado
		userconsultada := <-userResultado

		close(categoriaResultado)
		close(currencyResultado)
		close(userResultado)

		repositoryitem := implrepositoryItems.ItemRepositoryImpl{}
		itemModel := modelsItems.ItemsModel{
			Site:        registroArchivo[0],
			Id:          registroArchivo[1],
			Price:       itemConsultado.Body.Price,
			Start_time:  itemConsultado.Body.Date_created,
			Name:        categoriaConsultada.Name,
			Description: currencyConsultada.Description,
			Nickname:    userconsultada.Seller.Nickname,
		}
		createItem(repositoryitem, itemModel)
	}

}
