package middlewares

import (
	"net/http"

	"github.com/SpyroCode/tweeter/db"
)

//ChequeoBD es el middlewware que merite revisar la base de datos
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
