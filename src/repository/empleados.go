package repository

import "villas.com/src/modelos"

type Employ interface {
	SearchByDNI(dni string) (*modelos.Employ, error)             //ok
	SearchByName(nombre string) ([]*modelos.Employ, error)       //ok
	CantidadPorRegimen() ([]*modelos.EmployeesForRegimen, error) //ok
}
