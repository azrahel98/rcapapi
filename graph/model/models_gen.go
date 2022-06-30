// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Asistencia struct {
	Fecha *string `json:"fecha"`
	Dni   *string `json:"dni"`
	Hora  *string `json:"hora"`
	Hora2 *string `json:"hora2"`
	Hora3 *string `json:"hora3"`
	Reloj *string `json:"reloj"`
}

type Empleado struct {
	Dni     *string `json:"dni"`
	Nombre  string  `json:"nombre"`
	Ingreso *string `json:"ingreso"`
	Area    *string `json:"area"`
	Cargo   *string `json:"cargo"`
	Regimen *string `json:"regimen"`
	Horario *string `json:"horario"`
}

type EmployForRegimen struct {
	Regimen  string `json:"regimen"`
	Cantidad int    `json:"cantidad"`
}

type EmployI struct {
	Dni        *string `json:"dni"`
	Nombre     string  `json:"nombre"`
	Ingreso    *string `json:"ingreso"`
	Onomastico *string `json:"onomastico"`
	Cargo      *string `json:"cargo"`
	Sueldo     *string `json:"sueldo"`
}

type Papeleta struct {
	ID       *int       `json:"id"`
	Nombre   *string    `json:"nombre"`
	Fecha    *string    `json:"fecha"`
	Empleado *string    `json:"empleado"`
	Tipoper  *string    `json:"tipoper"`
	Descrip  *string    `json:"descrip"`
	Detalle  *string    `json:"detalle"`
	Retorno  *RetornoPa `json:"retorno"`
}

type PapeletaInput struct {
	Nombre   string           `json:"nombre"`
	Fecha    string           `json:"fecha"`
	Empleado string           `json:"empleado"`
	TipoP    PermisosPapeleta `json:"tipoP"`
	Descrip  string           `json:"descrip"`
	Detalle  string           `json:"detalle"`
	Retorno  RetornoPa        `json:"retorno"`
}

type Token struct {
	Token *string `json:"token"`
	Admin *string `json:"admin"`
}

type User struct {
	ID       *int    `json:"id"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Isadmin  *bool   `json:"isadmin"`
}

type PermisosPapeleta string

const (
	PermisosPapeletaDf          PermisosPapeleta = "DF"
	PermisosPapeletaAc          PermisosPapeleta = "AC"
	PermisosPapeletaJustificado PermisosPapeleta = "JUSTIFICADO"
	PermisosPapeletaOmision     PermisosPapeleta = "OMISION"
	PermisosPapeletaDfxhel      PermisosPapeleta = "DFXHEL"
	PermisosPapeletaOnomastico  PermisosPapeleta = "ONOMASTICO"
)

var AllPermisosPapeleta = []PermisosPapeleta{
	PermisosPapeletaDf,
	PermisosPapeletaAc,
	PermisosPapeletaJustificado,
	PermisosPapeletaOmision,
	PermisosPapeletaDfxhel,
	PermisosPapeletaOnomastico,
}

func (e PermisosPapeleta) IsValid() bool {
	switch e {
	case PermisosPapeletaDf, PermisosPapeletaAc, PermisosPapeletaJustificado, PermisosPapeletaOmision, PermisosPapeletaDfxhel, PermisosPapeletaOnomastico:
		return true
	}
	return false
}

func (e PermisosPapeleta) String() string {
	return string(e)
}

func (e *PermisosPapeleta) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PermisosPapeleta(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PermisosPapeleta", str)
	}
	return nil
}

func (e PermisosPapeleta) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RetornoPa string

const (
	RetornoPaY RetornoPa = "Y"
	RetornoPaN RetornoPa = "N"
)

var AllRetornoPa = []RetornoPa{
	RetornoPaY,
	RetornoPaN,
}

func (e RetornoPa) IsValid() bool {
	switch e {
	case RetornoPaY, RetornoPaN:
		return true
	}
	return false
}

func (e RetornoPa) String() string {
	return string(e)
}

func (e *RetornoPa) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RetornoPa(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RetornoPa", str)
	}
	return nil
}

func (e RetornoPa) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
