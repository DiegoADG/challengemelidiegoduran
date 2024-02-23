package impl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/DiegoADG/challengemelidiegoduran/integration/consultarCategoria/dto"
	"github.com/DiegoADG/challengemelidiegoduran/util"
)

type ConsultarCategoriaServiceImpl struct {
}

func (C ConsultarCategoriaServiceImpl) ConsultarCategorias(authorization string, idCategoria string) dto.ConsultarCategoriaDto {

	url := strings.Replace(util.URL_CATEGORIAS, "$CATEGORY_ID", idCategoria, -1)

	request, error := http.NewRequest(http.MethodGet, url, nil)
	if error != nil {
		fmt.Println("error request" + error.Error())
		return dto.ConsultarCategoriaDto{}
	}
	request.Header.Add("Authorization", authorization)

	response, error := http.DefaultClient.Do(request)
	if error != nil {
		fmt.Println("error response" + error.Error())
		return dto.ConsultarCategoriaDto{}
	}

	fmt.Printf("respuuesta: " + response.Status)

	responseBytes, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println("error response" + error.Error())
		return dto.ConsultarCategoriaDto{}
	}
	var categoria dto.ConsultarCategoriaDto

	if error := json.Unmarshal(responseBytes, &categoria); error != nil {
		fmt.Println("error unmarshal" + error.Error())
	}

	fmt.Printf("respuuesta: " + categoria.Name)
	return categoria
}
