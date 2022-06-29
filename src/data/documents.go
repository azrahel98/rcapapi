package data

import (
	"villas.com/graph/model"
	"villas.com/src/modelos"
	"villas.com/src/repository"
)

type DocumentsData struct {
	impl repository.Documents
}

func NewDocumentsData(r repository.Documents) DocumentsData {
	return DocumentsData{impl: r}
}

func (d DocumentsData) CrearPapeleta(e model.PapeletaInput) (*model.Papeleta, error) {
	doc, err := d.impl.Crear_Papeleta(modelos.Papeleta{
		Nombre:  e.Nombre,
		Fecha:   e.Fecha,
		Dni:     e.Empleado,
		Permiso: string(e.TipoP),
		Descrip: e.Descrip,
		Detalle: e.Detalle,
		Retorno: e.Retorno.String(),
	})
	if err != nil {
		return nil, err
	}
	return &model.Papeleta{
		ID:       &doc.Id,
		Nombre:   &doc.Nombre,
		Fecha:    &doc.Fecha,
		Empleado: &doc.Dni,
		Tipoper:  &doc.Permiso,
		Descrip:  &doc.Permiso,
		Detalle:  &doc.Detalle,
		Retorno:  (*model.RetornoPa)(&doc.Retorno),
	}, err
}

func (d DocumentsData) BuscarDocumentosPorDNI(dni string) (*model.Papeleta, error) {
	p, err := d.impl.FindByDniPapeletas(dni)
	if err != nil {
		return nil, err
	}
	return &model.Papeleta{
		ID:       &p.Id,
		Nombre:   &p.Nombre,
		Fecha:    &p.Fecha,
		Empleado: &p.Dni,
		Tipoper:  &p.Permiso,
		Descrip:  &p.Descrip,
		Detalle:  &p.Detalle,
		Retorno:  (*model.RetornoPa)(&p.Retorno),
	}, err

}

// func (d DocumentsData) UltimosRegistros() ([]*modelos.DocHistory, error) {
// 	doc, err := d.impl.Historial()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return doc, nil
// }

// func (d DocumentsData) DocumentosRecibidosxMes(mes int) ([]*modelos.DocsRecibidos, error) {
// 	docs, err := d.impl.DocumentosRecibidosxMes(mes, 2022)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return docs, nil
// }
// func (d DocumentsData) DocumentoPorID(id string) (*modelos.DocsRecibidos, error) {
// 	docs, err := d.impl.DocumentosRecibidorxID(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return docs, nil
// }

// func (d DocumentsData) EliminarporId(id int) error {
// 	err := d.impl.DeleteForID(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (d DocumentsData) DocumentosporMesyDNi(mes int, dni string) ([]*modelos.DocsRecibidos, error) {
// 	data, err := d.impl.DocumentsxMesyDni(mes, dni)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil

// }

// func (d DocumentsData) ActualizarDocum(doc *model.DocInpt) (*model.Document, error) {

// 	dx, err := d.impl.UpdateDoc(modelos.Documents{
// 		Id:       *doc.ID,
// 		Nombre:   doc.Nombre,
// 		Fecha:    doc.Fecha,
// 		Empleado: doc.Empleado,
// 		TipoDoc:  doc.Tipod.String(),
// 		TipoPer:  doc.TipoP.String(),
// 		Descrip:  doc.Descrip,
// 		RefDoc:   doc.Refdoc,
// 		Detalle:  doc.Detalle,
// 		Inicio:   doc.Inicio,
// 		Fin:      doc.Fin,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &model.Document{
// 		ID:       &dx.Id,
// 		Nombre:   &dx.Nombre,
// 		Fecha:    &dx.Fecha,
// 		Empleado: &dx.Empleado,
// 		Tipodoc:  &dx.TipoDoc,
// 		Tipoper:  &dx.TipoPer,
// 		Descrip:  &dx.Descrip,
// 		Refdoc:   &dx.RefDoc,
// 		Detalle:  &dx.Detalle,
// 		Inicio:   &dx.Inicio,
// 		Fin:      &dx.Fin,
// 	}, nil
// }
