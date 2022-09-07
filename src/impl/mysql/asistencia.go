package mysql

import (
	"database/sql"
	"strconv"

	"villas.com/src/modelos"
	"villas.com/src/service"
)

type AsistenciaImpl struct{}

func (a AsistenciaImpl) Buscar_Asistencia(dni string, mes int) ([]*modelos.Asistencia, error) {
	db, _ := service.GetInstance()
	var data []*modelos.Asistencia
	err := db.Query("select IFNULL(dni,'NA') dni ,IFNULL(nombre,'NA')nombre ,IFNULL(LEFT(hora,5),'NA') hora ,IFNULL(fecha,'NA') , IFNULL(reloj,'NA') reloj  from Asiste a where dni = ? and MONTH (fecha) = ?",
		func(r *sql.Rows) error {
			for r.Next() {
				var a modelos.Asistencia
				err := r.Scan(&a.Dni, &a.Nombre, &a.Hora, &a.Fecha, &a.Equipo)
				if err != nil {
					return err
				}
				data = append(data, &a)
			}
			return nil
		}, dni, strconv.Itoa(mes))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a AsistenciaImpl) Token_Check(dni string, token modelos.Token) error {
	db, _ := service.GetInstance()
	_, err := db.Exec("INSERT INTO Tokens (token,fend) values(?,?)", token.Value, token.FechaV)
	if err != nil {
		return err
	}
	return nil
}
