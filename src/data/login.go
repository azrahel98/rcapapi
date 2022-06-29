package data

import (
	"errors"

	"villas.com/src/modelos"
	"villas.com/src/repository"
)

type LoginData struct {
	re repository.LoginRepository
}

func NewLoginData(e repository.LoginRepository) LoginData {
	return LoginData{re: e}
}

func (e LoginData) Login(username, password string) (*modelos.User, error) {
	user, err := e.re.Login(username, password)
	if err != nil {
		return nil, errors.New("usuario no existe")
	}
	if user.Password != password {
		return nil, errors.New("contraseña incorrecta")
	}
	return user, nil
}
