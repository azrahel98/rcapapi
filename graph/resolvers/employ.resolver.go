package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"villas.com/graph/model"
	"villas.com/middleware"
	"villas.com/src/data"
	"villas.com/src/impl/mysql"
	"villas.com/src/modelos"
)

func (r *mutationResolver) CrearEmpleado(ctx context.Context, input *model.EmployI) (*model.Empleado, error) {
	err := middleware.GraphQlErrorHandler(ctx, false)
	if err != nil {
		return nil, err
	}
	employ := data.NewEmployData(mysql.EmpleadoImpl{})
	em, err := employ.CreateEmploy(modelos.Employ{
		Dni:     *input.Dni,
		Nombre:  input.Nombre,
		Ingreso: *input.Ingreso,

		Area:  "",
		Cargo: *input.Cargo,

		Regimen: "",
		Horario: "",
	})
	if err != nil {
		return nil, err
	}
	return &model.Empleado{
		Dni:     &em.Dni,
		Nombre:  em.Nombre,
		Ingreso: &em.Ingreso,

		Area:  &em.Area,
		Cargo: &em.Cargo,

		Regimen: &em.Regimen,
	}, nil
}

func (r *mutationResolver) BorrarEmpleado(ctx context.Context, dni *string) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) EmpleadoByDni(ctx context.Context, dni *string) (*model.Empleado, error) {
	err := middleware.GraphQlErrorHandler(ctx, false)
	if err != nil {
		return nil, err
	}
	err = middleware.GraphQlErrorHandler(ctx, false)
	if err != nil {
		return nil, err
	}
	employ := data.NewEmployData(mysql.EmpleadoImpl{})
	em, err := employ.FindByID(*dni)
	if err != nil {
		return nil, err
	}
	return &model.Empleado{
		Dni:     &em.Dni,
		Nombre:  em.Nombre,
		Ingreso: &em.Ingreso,
		Area:    &em.Area,
		Cargo:   &em.Cargo,
		Regimen: &em.Regimen,
	}, nil
}

func (r *queryResolver) EmpleadosRegimen(ctx context.Context) ([]*model.EmployForRegimen, error) {
	err := middleware.GraphQlErrorHandler(ctx, false)
	if err != nil {
		return nil, err
	}
	employ := data.NewEmployData(mysql.EmpleadoImpl{})
	res, err := employ.TrabajadoresxRegimen()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *queryResolver) EmpleadosSearch(ctx context.Context, nombre *string) ([]*model.Empleado, error) {
	err := middleware.GraphQlErrorHandler(ctx, false)
	if err != nil {
		return nil, err
	}
	if nombre == nil {
		return nil, errors.New("campo requerido")
	}
	employ := data.NewEmployData(mysql.EmpleadoImpl{})
	result, err := employ.BuscarEmpleado(*nombre)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Empleado(ctx context.Context, dni *string) (*model.Empleado, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Saludo(ctx context.Context) (*string, error) {

	saludo := "hola mundo"
	return &saludo, nil
}
