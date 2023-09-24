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

//AñadirEmpleado-EliminarEmpleado-OrdenarEmpleado/PorApellido

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
