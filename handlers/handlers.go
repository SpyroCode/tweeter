package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/SpyroCode/tweeter/middlewares"
	"github.com/SpyroCode/tweeter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores seteo mi puerto , el handlear y pongo a escucha mi servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HanderFunc("/registro", middlewares.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8088"
	}
	handlers := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handlers))
}
