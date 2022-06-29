package modelos

type Asistencia struct {
	Dni    string
	Nombre string
	Hora   string
	Fecha  string
	Equipo string
}

type Token struct {
	Value  string
	FechaV string
}

type DecodeToken struct {
	Dni string
	Mes int
}
