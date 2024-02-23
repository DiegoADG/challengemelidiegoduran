package consultarCurrency

import (
	"github.com/DiegoADG/challengemelidiegoduran/integration/consultarCurrency/dto"
)

type ConsultarCurrencyService interface {
	ConsultarCurrencys(authorization string, idCurrency string) dto.ConsultarCurrencyDto
}
