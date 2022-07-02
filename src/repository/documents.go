package repository

import "villas.com/src/modelos"

type Documents interface {
	Crear_Papeleta(employ modelos.Papeleta) (*modelos.Papeleta, error)
	FindByDniPapeletas(dni string) (*modelos.Papeleta, error)
	Create_Doc(e modelos.Docs, rango bool) (*modelos.Docs, error)
}
