package repository

import "villas.com/src/modelos"

type Employ interface {
	SearchByDNI(dni string) (*modelos.Employ, error)             //ok
	Todos() (*[]modelos.Employ, error)                           //pendiente
	Crear(employ modelos.Employ) (*modelos.Employ, error)        // pendiente
	SearchByName(nombre string) ([]*modelos.Employ, error)       //ok
	CantidadPorRegimen() ([]*modelos.EmployeesForRegimen, error) //ok
}
