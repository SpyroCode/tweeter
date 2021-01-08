package main

import (
	"log"

	"github.com/SpyroCode/tweeter/db"
	"github.com/SpyroCode/tweeter/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin Cnexion a la BD")
		return
	}

	handlers.Manejadores()

}
