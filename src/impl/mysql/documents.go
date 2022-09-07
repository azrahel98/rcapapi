package mysql

import (
	"database/sql"
	"errors"
	"log"

	"github.com/go-sql-driver/mysql"
	"villas.com/src/modelos"
	"villas.com/src/service"
)

type DocumentsImpl struct{}

// func (d DocumentsImpl) SearchById(id int) (*modelos.Documents, error) {
// 	var doc modelos.Documents
// 	db, _ := service.GetInstance()

// 	err := db.Query("select id,nombre,fecha,empleado,tipod,tipoPer,descrip,refDoc,detalle,inicio,fin from Documentos where id = ?", func(r *sql.Rows) error {
// 		for r.Next() {
// 			var ref, inicio, fin sql.NullString

// 			err := r.Scan(&doc.Id, &doc.Nombre, &doc.Fecha, &doc.Empleado, &doc.TipoDoc, &doc.TipoPer, &doc.Descrip, &ref, &doc.Detalle, &inicio, &fin)
// 			if err != nil {
// 				log.Println(err)
// 				return err
// 			}
// 			doc.RefDoc = ref.String
// 			doc.Inicio = ref.String
// 			doc.Fin = ref.String
// 		}
// 		return nil
// 	}, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Println(doc)
// 	return &doc, nil
// }

func (d DocumentsImpl) FindByDniPapeletas(dni string, month int) ([]*modelos.Papeleta, error) {
	db, _ := service.GetInstance()

	var pa []*modelos.Papeleta

	err := db.Query("select papeleta ,fecha ,permiso ,descr ,detalle  from Papeleta p where dni = ? and month(fecha) = ?", func(r *sql.Rows) error {
		for r.Next() {
			var p modelos.Papeleta
			err := r.Scan(&p.Nombre, &p.Fecha, &p.Permiso, &p.Descrip, &p.Detalle)
			if err != nil {
				return err
			}
			pa = append(pa, &p)
		}
		return nil
	}, dni, month)
	if err != nil {
		return nil, err
	}
	return pa, nil
}

func (d DocumentsImpl) Crear_Papeleta(e modelos.Papeleta) (*modelos.Papeleta, error) {
	db, _ := service.GetInstance()

	var id int64

	re, err := db.Exec("insert into Papeleta(dni,papeleta,descr,detalle,retorno,fecha,permiso) values(?,?,?,?,?,?,?)",
		e.Dni, e.Nombre, e.Descrip, e.Detalle, e.Retorno, e.Fecha, e.Permiso)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	id = *re

	return &modelos.Papeleta{
		Id:      int(id),
		Nombre:  e.Nombre,
		Fecha:   e.Fecha,
		Dni:     e.Dni,
		Permiso: e.Permiso,
		Descrip: e.Descrip,
		Detalle: e.Detalle,
		Retorno: e.Retorno,
	}, nil
}

func (d DocumentsImpl) Create_Doc(e modelos.Docs, rango bool) (*modelos.Docs, error) {
	db, _ := service.GetInstance()
	query := ""
	var id int64
	if rango {
		query = "insert into Doc(dni,doc,fecha,TipoDoc,tipoper,descrip,ref,inicio,fin) values(?,?,?,?,?,?,?,?,?)"
		res, err := db.Exec(query, e.Dni, e.Doc, e.Fecha, e.Tipo, e.Permi, e.Descrip, e.Ref, e.Inicio, e.Fin)
		if err != nil {
			me, _ := err.(*mysql.MySQLError)
			if me.Number == 1062 {
				return nil, errors.New("el documento ya existe")
			}
			return nil, err
		}
		id = *res

	} else {
		query = "insert into Doc(dni,doc,fecha,TipoDoc,tipoper,descrip,ref) values(?,?,?,?,?,?,?)"
		res, err := db.Exec(query, e.Dni, e.Doc, e.Fecha, e.Tipo, e.Permi, e.Descrip, e.Ref)
		if err != nil {
			me, _ := err.(*mysql.MySQLError)
			if me.Number == 1062 {
				return nil, errors.New("el documento ya existe")
			}
			return nil, err
		}
		id = *res

	}

	e.Id = int(id)
	return &e, nil

}

func (d DocumentsImpl) DocsByDNI(dni string, month int) ([]*modelos.Docs, error) {
	db, _ := service.GetInstance()

	var docs []*modelos.Docs

	err := db.Query("select docId ,dni,doc,fecha ,TipoDoc ,tipoper ,IFNULL(descrip,'') descrip ,IFNULL(`ref`,'') refer, IFNULL(inicio,'0000-00-00') inicio ,IFNULL(fin,'0000-00-00') fin  from Doc d where MONTH(fecha) = ? and dni = ?", func(r *sql.Rows) error {
		for r.Next() {
			var d modelos.Docs
			err := r.Scan(&d.Id, &d.Dni, &d.Doc, &d.Fecha, &d.Tipo, &d.Permi, &d.Descrip, &d.Ref, &d.Inicio, &d.Fin)
			if err != nil {
				return err
			}
			docs = append(docs, &d)
		}
		return nil
	}, month, dni)
	if err != nil {
		return nil, err
	}
	if len(docs) <= 0 {
		return make([]*modelos.Docs, 0), nil
	}
	return docs, nil
}
func (d DocumentsImpl) VacacionesByDNI(dni string, month int) ([]*modelos.Docs, error) {
	db, _ := service.GetInstance()

	var docs []*modelos.Docs

	err := db.Query("select docId ,dni,doc,fecha ,TipoDoc ,tipoper ,IFNULL(descrip,'') descrip ,IFNULL(`ref`,'') refer, IFNULL(inicio,'0000-00-00') inicio ,IFNULL(fin,'0000-00-00') fin  from Doc d where MONTH(inicio) <= ? and MONTH(fin) >= ? and d.dni  = ?", func(r *sql.Rows) error {
		for r.Next() {
			var d modelos.Docs
			err := r.Scan(&d.Id, &d.Dni, &d.Doc, &d.Fecha, &d.Tipo, &d.Permi, &d.Descrip, &d.Ref, &d.Inicio, &d.Fin)
			if err != nil {
				return err
			}
			docs = append(docs, &d)
		}
		return nil
	}, month, month, dni)
	if err != nil {
		return nil, err
	}
	if len(docs) <= 0 {
		return make([]*modelos.Docs, 0), nil
	}
	return docs, nil
}

func (d DocumentsImpl) Update_Papeleta(p modelos.Papeleta) (*modelos.Papeleta, error) {
	db, _ := service.GetInstance()

	_, err := db.Exec("UPDATE rcap.Papeleta SET papeleta=?,descr=?,detalle=?,fecha=?,permiso=? WHERE papeleta=?",
		p.Nombre, p.Descrip, p.Detalle, p.Fecha, p.Permiso, p.Nombre)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// func (d DocumentsImpl) Historial() ([]*modelos.DocHistory, error) {
// 	db, err := service.GetInstance()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var res []*modelos.DocHistory
// 	err = db.Query("SELECT * FROM historial", func(r *sql.Rows) error {
// 		for r.Next() {
// 			var doc modelos.DocHistory
// 			err := r.Scan(&doc.Id, &doc.TipoDo, &doc.Nombre, &doc.Fecha, &doc.Descrip, &doc.Empleado, &doc.Cargo)
// 			if err != nil {
// 				return err
// 			}
// 			res = append(res, &doc)
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

// func (d DocumentsImpl) DocumentosRecibidosxMes(mes int, año int) ([]*modelos.DocsRecibidos, error) {
// 	db, err := service.GetInstance()
// 	if err != nil {
// 		return nil, err
// 	}

// 	docs, err := db.DocsxMes(mes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return docs, nil

// }

// func (d DocumentsImpl) DocumentosRecibidorxID(document string) (*modelos.DocsRecibidos, error) {
// 	db, err := service.GetInstance()
// 	if err != nil {
// 		return nil, err
// 	}

// 	doc, err := db.DocsxMesxName(document)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return doc, err

// }

// func (d DocumentsImpl) UpdateDoc(docx modelos.Documents) (*modelos.Documents, error) {
// 	db, err := service.GetInstance()
// 	if err != nil {
// 		return nil, err
// 	}

// 	_, err = db.Exec("update `Documentos` set nombre = ?,fecha = ?,empleado = ?,tipod =?,tipoPer = ?,descrip = ?,refDoc = ?,detalle = ? where id = ?",
// 		docx.Nombre, docx.Fecha, docx.Empleado, docx.TipoDoc, docx.TipoPer, docx.Descrip, docx.RefDoc, docx.Detalle, docx.Id)
// 	if err != nil {
// 		log.Println("Aqui")
// 		return nil, err
// 	}
// 	return &docx, nil
// }

// func (d DocumentsImpl) DeleteForID(id int) error {
// 	db, err := service.GetInstance()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = db.Exec("delete from Documentos where id = ?", id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (d DocumentsImpl) DocumentsxMesyDni(mes int, dni string) ([]*modelos.DocsRecibidos, error) {
// 	db, err := service.GetInstance()
// 	if err != nil {
// 		return nil, err
// 	}
// 	var docs []*modelos.DocsRecibidos
// 	err = db.Query("SELECT * FROM documentosrecibidos where empleado = ? and month(fecha) = ?;", func(r *sql.Rows) error {
// 		for r.Next() {
// 			var d modelos.DocsRecibidos
// 			err := r.Scan(&d.Id, &d.TipDoc, &d.Nombre, &d.Empleado, &d.Dni, &d.Fecha, &d.TipoPer, &d.Descrip, &d.Detalle, &d.RefDoc, &d.Inicio, &d.Fin, &d.Create)
// 			if err != nil {
// 				return err
// 			}
// 			docs = append(docs, &d)
// 		}
// 		return nil
// 	}, dni, mes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return docs, nil
// }
