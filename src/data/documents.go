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

func (d DocumentsData) BuscarDocumentosPorDNI(dni string, month int) ([]*model.Papeleta, error) {
	p, err := d.impl.FindByDniPapeletas(dni, month)
	if err != nil {
		return nil, err
	}
	var r []*model.Papeleta
	for _, v := range p {
		r = append(r, &model.Papeleta{
			ID:       &v.Id,
			Nombre:   &v.Nombre,
			Fecha:    &v.Fecha,
			Empleado: &v.Dni,
			Tipoper:  &v.Permiso,
			Descrip:  &v.Descrip,
			Detalle:  &v.Detalle,
			Retorno:  (*model.RetornoPa)(&v.Retorno),
		})
	}

	return r, nil
}
func (d DocumentsData) BuscarDocsByDni(dni string, month int) ([]*model.Docs, error) {
	var docs []*model.Docs
	p, err := d.impl.VacacionesByDNI(dni, month)
	if err != nil {
		return nil, err
	}
	r, err := d.impl.DocsByDNI(dni, month)
	if err != nil {
		return nil, err
	}
	for _, x := range r {
		docs = append(docs, &model.Docs{
			ID:      &x.Id,
			Dni:     &x.Dni,
			Doc:     &x.Doc,
			Fecha:   &x.Fecha,
			Descrip: &x.Descrip,
			Ref:     &x.Ref,
			Inicio:  &x.Inicio,
			Fin:     &x.Fin,
			Tipo:    (*model.TiposDocs)(&x.Tipo),
			Permiso: (*model.PermisosDoc)(&x.Permi),
		})
	}

	if len(docs) > 0 {
		for _, v := range p {
			for _, x := range docs {
				if v.Doc != *x.Doc {
					docs = append(docs, &model.Docs{
						ID:      &v.Id,
						Dni:     &v.Dni,
						Doc:     &v.Doc,
						Fecha:   &v.Fecha,
						Descrip: &v.Descrip,
						Ref:     &v.Ref,
						Inicio:  &v.Inicio,
						Fin:     &v.Fin,
						Tipo:    (*model.TiposDocs)(&v.Tipo),
						Permiso: (*model.PermisosDoc)(&v.Permi),
					})
				}
			}

		}
	} else {
		for _, x := range p {
			docs = append(docs, &model.Docs{
				ID:      &x.Id,
				Dni:     &x.Dni,
				Doc:     &x.Doc,
				Fecha:   &x.Fecha,
				Descrip: &x.Descrip,
				Ref:     &x.Ref,
				Inicio:  &x.Inicio,
				Fin:     &x.Fin,
				Tipo:    (*model.TiposDocs)(&x.Tipo),
				Permiso: (*model.PermisosDoc)(&x.Permi),
			})
		}
	}

	return unique(docs), nil
}
func (d DocumentsData) Update_Papeleta(papeleta model.PapeletaInput) (*model.Papeleta, error) {

	res, err := d.impl.Update_Papeleta(modelos.Papeleta{
		Nombre:  papeleta.Nombre,
		Fecha:   papeleta.Fecha,
		Permiso: string(papeleta.TipoP),
		Descrip: papeleta.Descrip,
		Detalle: papeleta.Detalle,
	})
	if err != nil {
		return nil, err
	}
	return &model.Papeleta{
		ID:       &res.Id,
		Nombre:   &res.Nombre,
		Fecha:    &res.Fecha,
		Empleado: &res.Dni,
		Tipoper:  &res.Permiso,
		Descrip:  &res.Descrip,
		Detalle:  &res.Detalle,
		Retorno:  (*model.RetornoPa)(&res.Retorno),
	}, nil
}

func unique(s []*model.Docs) []*model.Docs {
	inResult := make(map[*model.Docs]bool)
	var result []*model.Docs
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
