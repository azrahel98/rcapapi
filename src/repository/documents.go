package repository

import "villas.com/src/modelos"

type Documents interface {
	Crear_Papeleta(employ modelos.Papeleta) (*modelos.Papeleta, error)
	Create_Doc(e modelos.Docs, rango bool) (*modelos.Docs, error)
	Update_Papeleta(p modelos.Papeleta) (*modelos.Papeleta, error)
	FindByDniPapeletas(dni string, month int) ([]*modelos.Papeleta, error)
	DocsByDNI(dni string, month int) ([]*modelos.Docs, error)

	VacacionesByDNI(dni string, month int) ([]*modelos.Docs, error)
}
