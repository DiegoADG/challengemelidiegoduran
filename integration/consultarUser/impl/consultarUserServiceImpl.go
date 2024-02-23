package impl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/DiegoADG/challengemelidiegoduran/integration/consultarUser/dto"
	"github.com/DiegoADG/challengemelidiegoduran/util"
)

type ConsultarUserServiceImpl struct {
}

func (C ConsultarUserServiceImpl) ConsultarUsers(authorization string, idSeller string, idSite string) dto.ConsultarUserDto {

	url := strings.Replace(util.URL_USERS, "$SITE_ID", idSite, -1)
	url = strings.Replace(url, "$SELLER_ID", idSeller, -1)
	fmt.Println(url)
	request, error := http.NewRequest(http.MethodGet, url, nil)
	if error != nil {
		fmt.Println("error request" + error.Error())
		return dto.ConsultarUserDto{}
	}
	request.Header.Add("Authorization", authorization)

	response, error := http.DefaultClient.Do(request)
	if error != nil {
		fmt.Println("error response" + error.Error())
		return dto.ConsultarUserDto{}
	}

	fmt.Printf("respuuesta: " + response.Status)

	responseBytes, error := io.ReadAll(response.Body)

	if error != nil {
		fmt.Println("error response" + error.Error())
		return dto.ConsultarUserDto{}
	}
	var user dto.ConsultarUserDto

	if error := json.Unmarshal(responseBytes, &user); error != nil {
		fmt.Println("error unmarshal" + error.Error())
	}

	fmt.Printf("respuuesta: " + user.Seller.Nickname)
	return user
}
