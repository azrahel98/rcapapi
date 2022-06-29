package repository

import "villas.com/src/modelos"

type LoginRepository interface {
	Login(username, pass string) (*modelos.User, error)
}
