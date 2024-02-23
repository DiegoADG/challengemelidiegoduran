package impl

import (
	"context"

	"github.com/DiegoADG/challengemelidiegoduran/config"
	"github.com/DiegoADG/challengemelidiegoduran/domain/consultarItems/dto"
	"github.com/DiegoADG/challengemelidiegoduran/persistence/models"
	"go.mongodb.org/mongo-driver/bson"
)

type ItemRepositoryImpl struct {
}

var client = config.GetConexion()
var ctx = context.Background()

func (i ItemRepositoryImpl) Create(item models.ItemsModel) error {

	var erro error
	var items bson.M
	erro = client.Database("challengemeli").Collection("items").FindOne(ctx, bson.D{{"id", item.Id}, {"site", item.Site}}).Decode(&items)

	if len(items) == 0 {

		_, erro = client.Database("challengemeli").Collection("items").InsertOne(ctx, item)
		if erro != nil {
			return erro
		}
	} else {
		update := bson.M{
			"$set": bson.M{
				"price":       item.Price,
				"start_time":  item.Start_time,
				"name":        item.Name,
				"description": item.Description,
				"nickname":    item.Nickname,
			},
		}

		_, erro = client.Database("challengemeli").Collection("items").UpdateOne(ctx, bson.D{{"id", item.Id}, {"site", item.Site}}, update)
		if erro != nil {
			return erro
		}
	}

	return nil
}

func (i ItemRepositoryImpl) Consultar() (dto.RespuestadeItemsdto, error) {
	var respuestaItems dto.RespuestadeItemsdto
	cur, erro := client.Database("challengemeli").Collection("items").Find(ctx, bson.D{})
	if erro != nil {
		return respuestaItems, erro
	}

	for cur.Next(ctx) {
		var respuestaItem dto.RespuestaItemsdto
		erro = cur.Decode(&respuestaItem)

		if erro != nil {
			return respuestaItems, erro
		}

		respuestaItems = append(respuestaItems, &respuestaItem)
	}
	return respuestaItems, nil

}
