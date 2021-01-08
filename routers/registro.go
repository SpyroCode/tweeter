package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SpyroCode/tweeter/db"
	"github.com/SpyroCode/tweeter/models"
)

// Registro es la funcion para crear la bel regsistro de usuario en la base de datos
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido"+err.Error(), 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe specificar una contraseña de almenos  6 caracteres"+err.Error(), 400)
		return
	}

	_, encontrado, _ := db.CheckUserExiste(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe este usuario con este Email", 400)
		return
	}

	_, status, err := db.UserInsert(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al insertar el registro"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el resgitro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
