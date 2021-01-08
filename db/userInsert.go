package db

import (
	"context"
	"time"

	"github.com/SpyroCode/tweeter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserInsert(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("tweeter")
	col := db.Collection("users")
	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
