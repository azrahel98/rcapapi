package data

import (
	"errors"

	"villas.com/graph/model"
	"villas.com/src/modelos"
	"villas.com/src/repository"
)

type AsisteData struct {
	impl repository.AsistenciaRep
}

func NewAsisteData(r repository.AsistenciaRep) AsisteData {
	return AsisteData{impl: r}
}

func (im AsisteData) BuscarAsistencia(dni string, mes int) ([]*model.Asistencia, error) {

	d, err := im.impl.Buscar_Asistencia(dni, mes)
	if len(d) == 0 {
		return nil, errors.New("sin datos")
	}

	if err != nil {
		return nil, err
	}
	var result []*model.Asistencia

	for _, v := range d {
		result = append(result, &model.Asistencia{
			Fecha: &v.Fecha,
			Dni:   &v.Dni,
			Hora:  &v.Hora,
			Reloj: &v.Equipo,
		})
	}
	return result, nil
}
func (im AsisteData) CrearTokenForUser(tok modelos.Token) error {
	err := im.impl.Token_Check("", tok)
	if err != nil {
		return err
	}
	return nil
}
