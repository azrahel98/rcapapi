package modelos

type Employ struct {
	Dni     string
	Nombre  string
	Ingreso string
	Area    string
	Cargo   string
	Regimen string
	Horario string
}

type EmployeesForRegimen struct {
	Cantidad int
	Nombre   string
}
