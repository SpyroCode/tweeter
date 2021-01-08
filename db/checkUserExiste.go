package db

import (
	"context"
	"time"

	"github.com/SpyroCode/tweeter/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUserExiste(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("tweeter")
	col := db.Collection("users")
	condiction := bson.M{"email": email}

	var resultado models.Usuario
	err := col.FindOne(ctx, condiction).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
