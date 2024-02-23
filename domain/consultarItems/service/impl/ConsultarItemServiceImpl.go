package impl

import (
	"fmt"

	dto "github.com/DiegoADG/challengemelidiegoduran/domain/consultarItems/dto"
	repositoryItems "github.com/DiegoADG/challengemelidiegoduran/persistence/repository"
	implrepositoryItems "github.com/DiegoADG/challengemelidiegoduran/persistence/repository/impl"
)

type ConsultarItemsServiceImpl struct {
}

func (C ConsultarItemsServiceImpl) ConsultarItems() dto.RespuestadeItemsdto {
	repositoryitem := implrepositoryItems.ItemRepositoryImpl{}
	return consultarItem(repositoryitem)
}
func consultarItem(repository repositoryItems.ItemRepository) dto.RespuestadeItemsdto {

	items, err := repository.Consultar()
	if err != nil {
		fmt.Println(err)
	}
	return items
}
