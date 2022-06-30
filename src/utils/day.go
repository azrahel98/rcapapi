package utils

import (
	"log"

	"villas.com/graph/model"
)

func OrdenarDias(data []*model.Asistencia) []*model.Asistencia {

	var result []*model.Asistencia

	for i := 0; i < len(data); i++ {
		if len(data)-1 > i {
			if *data[i].Fecha == *data[i+1].Fecha {

				if len(data)-2 > i {

					if *data[i].Fecha == *data[i+1].Fecha && *data[i].Fecha == *data[i+2].Fecha {
						log.Println(*data[i].Fecha)
						result = append(result, &model.Asistencia{
							Fecha: data[i].Fecha,
							Dni:   data[i].Dni,
							Hora:  data[i].Hora,
							Hora2: data[i+1].Hora,
							Hora3: data[i+2].Hora,
							Reloj: data[i].Reloj,
						})

					} else {
						result = append(result, &model.Asistencia{
							Fecha: data[i].Fecha,
							Dni:   data[i].Dni,
							Hora:  data[i].Hora,
							Hora2: data[i+1].Hora,
							Reloj: data[i].Reloj,
						})

					}
				}
			}
		}
	}
	return result
}
