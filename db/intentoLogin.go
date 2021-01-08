package db

import (
	"github.com/SpyroCode/tweeter/models"
	"golang.org/x/crypto/bcrypt"
)

// IntentoLogin realiza chequeo de login de lña BD
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := CheckUserExiste(email)
	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usu, false
	}

	return usu, true

}
