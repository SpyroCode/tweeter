package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN objeto de conexion*/
var MongoCN = ConectarBD()
var clienteOption = options.Client().ApplyURI("mongodb+srv://dbUser:8ZeZPRpLvvKwkds@cluster0.qtobi.mongodb.net/tweeter?retryWrites=true&w=majority")

/*ConectarBD es la funcion para conectar*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOption)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la DB")
	return client
}

/*CheckConnection es el ping a la base de datos*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
