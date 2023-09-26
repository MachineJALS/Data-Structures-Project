package empleado

import (
	"fmt"
	"time"
)

type DiaTrabajo struct {
	Fecha    time.Time
	Duracion time.Duration
}

type Trabajo struct {
	Tipo   string // "Jefe" o "Empleado"
	Sector string // "Cajero", "Bartender", "Informático"
	Jefe   *Empleado
}

type Empleado struct {
	Id           int
	Nombre       string
	Apellido     string
	Edad         int
	SalarioHora  float64
	TipoSalario  string   // "Mensual" o "Hora"
	Complementos []string // Lista de complementos
	Trabajo      Trabajo
	DiasTrabajo  []DiaTrabajo
}

type EmpleadosPorId []Empleado

func (a EmpleadosPorId) Len() int           { return len(a) }
func (a EmpleadosPorId) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a EmpleadosPorId) Less(i, j int) bool { return a[i].Id < a[j].Id } //sort.Sort(EmpleadosPorId(empleados))

type EmpleadosPorApellido []Empleado

func (a EmpleadosPorApellido) Len() int           { return len(a) }
func (a EmpleadosPorApellido) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a EmpleadosPorApellido) Less(i, j int) bool { return a[i].Apellido < a[j].Apellido } //sort.Sort(EmpleadosPorApellido(empleados))

func (e Empleado) CalcularSalario() (salario float64) {
	for _, dia := range e.DiasTrabajo {
		salario += dia.Duracion.Hours() * e.SalarioHora
	}
	return
}

func (e *Empleado) ObtenerJefe() *Empleado {
	return e.Trabajo.Jefe
}

func (e Empleado) String() string {
	return fmt.Sprintf("ID:%d,Nombre:%s,Apellido:%s,Edad:%d", e.Id, e.Nombre, e.Apellido, e.Edad)
}
