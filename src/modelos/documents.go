package modelos

type Papeleta struct {
	Id      int
	Nombre  string
	Fecha   string
	Dni     string
	Permiso string
	Descrip string
	Detalle string
	Retorno string
}

type Docs struct {
	Id      int
	Dni     string
	Doc     string
	Fecha   string
	Tipo    string
	Permi   string
	Descrip string
	Ref     string
	Inicio  string
	Fin     string
}
