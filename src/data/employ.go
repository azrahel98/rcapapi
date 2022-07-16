package data

import (
	"villas.com/graph/model"
	"villas.com/src/modelos"
	"villas.com/src/repository"
)

type EmployData struct {
	impl repository.Employ
}

func NewEmployData(r repository.Employ) EmployData {
	return EmployData{impl: r}
}

func (imp EmployData) BuscarEmpleado(nombre string) ([]*model.Empleado, error) {
	employ, err := imp.impl.SearchByName(nombre)
	if err != nil {
		return nil, err
	}
	var resultado []*model.Empleado
	for _, v := range employ {
		resultado = append(resultado, &model.Empleado{
			Dni:     &v.Dni,
			Nombre:  v.Nombre,
			Ingreso: &v.Ingreso,
			Area:    &v.Area,
			Cargo:   &v.Cargo,
			Regimen: &v.Regimen,
			Horario: &v.Horario,
		})
	}
	return resultado, nil
}

func (imp EmployData) FindByID(dni string) (*modelos.Employ, error) {
	employ, err := imp.impl.SearchByDNI(dni)
	if err != nil {
		return nil, err
	}
	return employ, nil
}

func (imp EmployData) TrabajadoresxRegimen() ([]*model.EmployForRegimen, error) {
	res, err := imp.impl.CantidadPorRegimen()
	if err != nil {
		return nil, err
	}
	var result []*model.EmployForRegimen

	for _, v := range res {
		result = append(result, &model.EmployForRegimen{
			Regimen:  v.Nombre,
			Cantidad: v.Cantidad,
		})
	}

	return result, err

}
