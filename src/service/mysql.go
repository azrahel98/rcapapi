package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlDB *Mysql

type Mysql struct {
	db *sql.DB
}

func GetInstance() (*Mysql, error) {
	if mysqlDB == nil {
		db, err := connectDB()
		if err != nil {
			return nil, err
		}
		mysqlDB = db
		return mysqlDB, nil
	}
	return mysqlDB, nil
}

func connectDB() (*Mysql, error) {

	db, err := sql.Open("mysql", os.Getenv("URIPATH"))
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("fallo la conexion a la base de datos")
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("esto es un error")
		log.Println(err)
	}
	return &Mysql{db: db}, nil
}

func (db *Mysql) Query(strq string, scan func(*sql.Rows) error, args ...any) error {
	query, err := db.db.Query(strq, args...)
	if err != nil {
		log.Println(err)
		return err
	}
	defer query.Close()
	return scan(query)

}

func (db *Mysql) Exec(strq string, args ...any) (*int64, error) {
	query, err := db.db.Exec(strq, args...)
	if err != nil {
		return nil, err
	}
	id, err := query.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}

// func (db *Mysql) DocsxMes(mes int) ([]*modelos.DocsRecibidos, error) {
// 	query, err := db.db.Query("SELECT * FROM documentosrecibidos where MONTH(fecha) = ? AND YEAR(fecha) = 2022", mes)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	defer query.Close()
// 	var docs []*modelos.DocsRecibidos
// 	for query.Next() {
// 		var d modelos.DocsRecibidos
// 		err := query.Scan(&d.Id, &d.TipDoc, &d.Nombre, &d.Empleado, &d.Dni, &d.Fecha, &d.TipoPer, &d.Descrip, &d.Detalle, &d.RefDoc, &d.Inicio, &d.Fin, &d.Create)
// 		if err != nil {
// 			log.Println(err)
// 			return nil, err
// 		}
// 		docs = append(docs, &d)
// 	}
// 	return docs, nil
// }

// func (db *Mysql) DocsxMesxName(id string) (*modelos.DocsRecibidos, error) {
// 	query, err := db.db.Query("SELECT * FROM documentosrecibidos where id = ?", id)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}
// 	defer query.Close()
// 	var docs *modelos.DocsRecibidos

// 	for query.Next() {
// 		var d modelos.DocsRecibidos
// 		err := query.Scan(&d.Id, &d.TipDoc, &d.Nombre, &d.Empleado, &d.Dni, &d.Fecha, &d.TipoPer, &d.Descrip, &d.Detalle, &d.RefDoc, &d.Inicio, &d.Fin, &d.Create)
// 		if err != nil {
// 			return nil, err
// 		}
// 		docs = &d
// 	}
// 	if docs == nil {
// 		return nil, errors.New("sin datos")
// 	}
// 	return docs, nil
// }
// func (db *Mysql) Asistencia(dni string, mes int) ([]*modelos.Asistencia, error) {
// 	var data []*modelos.Asistencia
// 	query, err := db.db.Query("select ifnull(a.cfecha,'-') as fecha,ifnull(a.hora,'-') as marca,ifnull(d.fecha,'-') as fechadoc,ifnull(d.nombre,'-') as nombre,ifnull(d.tipoPer,'-') as permiso,ifnull(d.descrip,'-') as descript,ifnull(d.detalle,'na') as detalle,ifnull(d.id,'na') as id from Documentos d RIGHT JOIN Asiste a on a.codsoc = d.empleado and a.cfecha = d.fecha where a.codsoc = ?  and month(cfecha) = ?", dni, mes)
// 	if err != nil {

// 		return nil, err
// 	}
// 	defer query.Close()
// 	for query.Next() {
// 		var a modelos.Asistencia
// 		err := query.Scan(&a.Fecha, &a.Marca, &a.FechaDoc, &a.Nombre, &a.Permiso, &a.Descrip, &a.Detalle, &a.ID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		data = append(data, &a)
// 	}
// 	return data, nil
// }
