package repository

import (
	"villas.com/src/modelos"
)

type AsistenciaRep interface {
	Buscar_Asistencia(dni string, mes int) ([]*modelos.Asistencia, error)
	Token_Check(dni string, token modelos.Token) error
}
