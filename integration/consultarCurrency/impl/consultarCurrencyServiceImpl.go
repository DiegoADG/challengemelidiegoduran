package impl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/DiegoADG/challengemelidiegoduran/integration/consultarCurrency/dto"
	"github.com/DiegoADG/challengemelidiegoduran/util"
)

type ConsultarCurrencyServiceImpl struct {
}

func (C ConsultarCurrencyServiceImpl) ConsultarCurrencys(authorization string, idCurrency string) dto.ConsultarCurrencyDto {

	url := strings.Replace(util.URL_CURRENCIES, "$CURRENCY_ID", idCurrency, -1)

	request, error := http.NewRequest(http.MethodGet, url, nil)
	if error != nil {
		fmt.Println("error request" + error.Error())
		return dto.ConsultarCurrencyDto{}
	}
	request.Header.Add("Authorization", authorization)

	response, error := http.DefaultClient.Do(request)
	if error != nil {
		fmt.Println("error response" + error.Error())
		return dto.ConsultarCurrencyDto{}
	}

	fmt.Printf("respuuesta: " + response.Status)

	responseBytes, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println("error response" + error.Error())
		return dto.ConsultarCurrencyDto{}
	}
	var currency dto.ConsultarCurrencyDto

	if error := json.Unmarshal(responseBytes, &currency); error != nil {
		fmt.Println("error unmarshal" + error.Error())
	}

	fmt.Printf("respuuesta: " + currency.Description)
	return currency
}
