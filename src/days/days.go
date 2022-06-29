package days

import (
	"fmt"
	"log"
	"sort"
	"time"

	"villas.com/src/modelos"
)

type ByDate []*modelos.Asistencia

func (a ByDate) Len() int { return len(a) }
func (a ByDate) Less(i, j int) bool {
	date, _ := time.Parse("2006-01-02T15:04:05Z", fmt.Sprintf("%sT%sZ", a[i].Fecha, a[i].Hora))
	date2, _ := time.Parse("2006-01-02T15:04:05Z", fmt.Sprintf("%sT%sZ", a[j].Fecha, a[j].Hora))
	return date.Before(date2)
}
func (a ByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func OrganizarDias(fechai time.Time, fechaf time.Time, asistencia []*modelos.Asistencia) []*modelos.Asistencia {

	mes := ""
	parseMes := fechai.UTC().Month()
	log.Println(fechai)
	if int(parseMes) < 10 {
		mes = fmt.Sprintf("0%d", int(parseMes))
	} else {
		mes = fmt.Sprintf("%d", int(parseMes))
	}

	iniciof := fmt.Sprintf("%d-%s-0%d", int(fechai.UTC().Year()), mes, int(fechai.UTC().Day()))
	finalf := fmt.Sprintf("%d-%s-%d", int(fechaf.UTC().Year()), mes, int(fechaf.UTC().Day()))

	if asistencia[0].Fecha != iniciof {
		asistencia = append(asistencia, &modelos.Asistencia{
			Fecha: iniciof,
			Hora:  "00:00:00",
		})
		sort.Sort(ByDate(asistencia))
	}
	if asistencia[len(asistencia)-1].Fecha != finalf {
		asistencia = append(asistencia, &modelos.Asistencia{
			Fecha: finalf,
			Hora:  "00:00:00",
		})
		sort.Sort(ByDate(asistencia))
	}

	for index, v := range asistencia {
		if index+1 < len(asistencia) {
			mañana, _ := time.Parse("2006-01-02", asistencia[index+1].Fecha)
			hoy, _ := time.Parse("2006-01-02", v.Fecha)
			diff := int(mañana.Day()) - int(hoy.Day())
			if diff > 1 {
				iter := 1
				for diff < iter {
					if int(hoy.UTC().Day())+iter < 10 {
						asistencia = append(asistencia, &modelos.Asistencia{
							Fecha: fmt.Sprintf("%d-%s-0%d", int(hoy.UTC().Year()), mes, int(hoy.UTC().Day())+iter),
							Hora:  "00:00:00",
						})
					} else {
						asistencia = append(asistencia, &modelos.Asistencia{
							Fecha: fmt.Sprintf("%d-%s-%d", int(hoy.UTC().Year()), mes, int(hoy.UTC().Day())+iter),
							Hora:  "00:00:00",
						})
					}
					iter++
				}

			}
		}
		sort.Sort(ByDate(asistencia))
	}

	sort.Sort(ByDate(asistencia))

	return asistencia
}
