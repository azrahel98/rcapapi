package mysql

import (
	"database/sql"
	"errors"
	"log"

	"villas.com/src/modelos"
	"villas.com/src/service"
)

type EmpleadoImpl struct{}

// func (e EmpleadoImpl) Todos() (*[]modelos.Employ, error) {
// 	db, err := service.GetInstance()
// 	if err != nil {
// 		return nil, err
// 	}
// 	empleados := []modelos.Employ{}
// 	err = db.Query("select * from Empleado", func(r *sql.Rows) error {
// 		for r.Next() {
// 			var e modelos.Employ
// 			err := r.Scan(&e.Dni, &e.Nombre, &e.Ingreso, &e.Area, &e.Cargo, &e.Regimen, &e.Horario)
// 			if err != nil {
// 				return err
// 			}
// 			empleados = append(empleados, e)
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &empleados, nil
// }

func (e EmpleadoImpl) SearchByDNI(dni string) (*modelos.Employ, error) {
	if dni == "" {
		return nil, errors.New("dni vacio")
	}
	var employ modelos.Employ

	db, _ := service.GetInstance()
	err := db.Query("select * from Detalle_Empleados de where dni = ?", func(r *sql.Rows) error {
		for r.Next() {
			err := r.Scan(&employ.Dni, &employ.Nombre, &employ.Ingreso, &employ.Regimen, &employ.Cargo, &employ.Area, &employ.Horario)
			if err != nil {
				return err
			}
		}
		if (employ == modelos.Employ{}) {
			return errors.New("el empleado no existe")
		}
		return nil
	}, dni)

	if err != nil {
		return nil, err
	}
	return &employ, nil
}

// func (e EmpleadoImpl) Crear(em modelos.Employ) (*modelos.Employ, error) {
// 	db, _ := service.GetInstance()
// 	if (em == modelos.Employ{}) {
// 		return nil, errors.New("campos Vacios")
// 	}
// 	err := db.Query("insert into Empleado(dni,nombre,ingreso,onomastico,cargo,sueldo) values(?,?,?,?,?,?)",
// 		func(r *sql.Rows) error { return nil }, em.Dni, em.Nombre, em.Ingreso, em.Cargo)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &modelos.Employ{
// 		Dni:     em.Dni,
// 		Nombre:  em.Nombre,
// 		Ingreso: em.Ingreso,
// 		Area:    em.Area,
// 		Cargo:   em.Cargo,
// 		Regimen: "CAS",
// 		Horario: "",
// 	}, nil
// }

func (e EmpleadoImpl) CantidadPorRegimen() ([]*modelos.EmployeesForRegimen, error) {
	db, _ := service.GetInstance()
	var resultado []*modelos.EmployeesForRegimen
	err := db.Query("select COUNT(e.regimen),r.nombre  from Employ e left join Regimen r on e.regimen = r.regid GROUP BY e.regimen ", func(r *sql.Rows) error {
		for r.Next() {
			var a modelos.EmployeesForRegimen
			err := r.Scan(&a.Cantidad, &a.Nombre)
			if err != nil {
				log.Println(err)
				return err
			}
			resultado = append(resultado, &a)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return resultado, nil
}
func (e EmpleadoImpl) SearchByName(nombre string) ([]*modelos.Employ, error) {
	db, _ := service.GetInstance()

	var employees []*modelos.Employ
	err := db.Query("SELECT * from Detalle_Empleados where nombre LIKE '%"+nombre+"%'", func(r *sql.Rows) error {
		for r.Next() {
			var e modelos.Employ
			err := r.Scan(&e.Dni, &e.Nombre, &e.Ingreso, &e.Regimen, &e.Cargo, &e.Area, &e.Horario)
			if err != nil {
				return err
			}
			employees = append(employees, &e)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return employees, nil
}
