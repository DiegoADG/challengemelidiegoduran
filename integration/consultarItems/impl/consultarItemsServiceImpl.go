package impl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/DiegoADG/challengemelidiegoduran/integration/consultarItems/dto"
	"github.com/DiegoADG/challengemelidiegoduran/util"
)

type ConsultarItemsServiceImpl struct {
}

func (i ConsultarItemsServiceImpl) ConsultarItems(authorization string, idItem string) dto.ConsultarItemsDto {

	url := strings.Replace(util.URL_ITEMS, "$ITEM_ID", idItem, -1)

	request, error := http.NewRequest(http.MethodGet, url, nil)
	if error != nil {
		fmt.Println("error request" + error.Error())
		return dto.ConsultarItemsDto{}
	}
	request.Header.Add("Authorization", authorization)

	response, error := http.DefaultClient.Do(request)
	if error != nil {
		fmt.Println("error response" + error.Error())
		return dto.ConsultarItemsDto{}
	}

	fmt.Printf("respuuesta: " + response.Status)

	responseBytes, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println("error response" + error.Error())
		return dto.ConsultarItemsDto{}
	}
	var items []dto.ConsultarItemsDto

	if error := json.Unmarshal(responseBytes, &items); error != nil {
		fmt.Println("error unmarshal" + error.Error())
	}

	fmt.Printf("respuuesta: " + items[0].Body.Category_id)
	return items[0]
}
