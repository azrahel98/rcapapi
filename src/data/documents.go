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

func (d DocumentsData) CrearDoc(e *model.DocInput, rang bool) (*model.Docs, error) {
	if rang {
		doc, err := d.impl.Create_Doc(modelos.Docs{
			Dni:     *e.Dni,
			Doc:     *e.Doc,
			Fecha:   *e.Fecha,
			Tipo:    e.Tipo.String(),
			Permi:   e.Permiso.String(),
			Descrip: *e.Descrip,
			Ref:     *e.Ref,
			Inicio:  *e.Inicio,
			Fin:     *e.Fin,
		}, true)
		if err != nil {
			return nil, err
		}
		return &model.Docs{
			ID:      &doc.Id,
			Dni:     &doc.Dni,
			Doc:     &doc.Doc,
			Fecha:   &doc.Fecha,
			Tipo:    (*model.TiposDocs)(&doc.Tipo),
			Permiso: (*model.PermisosDoc)(&doc.Permi),
			Descrip: &doc.Descrip,
			Ref:     &doc.Ref,
			Inicio:  &doc.Inicio,
			Fin:     &doc.Fin,
		}, err
	} else {
		doc, err := d.impl.Create_Doc(modelos.Docs{
			Dni:     *e.Dni,
			Doc:     *e.Doc,
			Fecha:   *e.Fecha,
			Tipo:    e.Tipo.String(),
			Permi:   e.Permiso.String(),
			Descrip: *e.Descrip,
			Ref:     *e.Ref,
		}, false)
		if err != nil {
			return nil, err
		}
		return &model.Docs{
			ID:      &doc.Id,
			Dni:     &doc.Dni,
			Doc:     &doc.Doc,
			Fecha:   &doc.Fecha,
			Tipo:    (*model.TiposDocs)(&doc.Tipo),
			Permiso: (*model.PermisosDoc)(&doc.Permi),
			Descrip: &doc.Descrip,
			Ref:     &doc.Ref,
		}, err
	}

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
