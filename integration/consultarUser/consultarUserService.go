package consultaruser

import (
	"github.com/DiegoADG/challengemelidiegoduran/integration/consultarUser/dto"
)

type ConsultarUserService interface {
	ConsultarUsers(authorization string, idSeller string, idSite string) dto.ConsultarUserDto
}
