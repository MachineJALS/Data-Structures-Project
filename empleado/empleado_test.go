package empleado

import (
	"testing"
	"time"
)

var emp1 = Empleado{
	Id:           1,
	Nombre:       "a",
	Apellido:     "L",
	Edad:         1,
	SalarioHora:  2,
	TipoSalario:  "Hora",
	Complementos: []string{},
	Trabajo:      Trabajo{},
	DiasTrabajo: []DiaTrabajo{{
		Fecha:    time.Time{},
		Duracion: 8 * time.Hour,
	}},
}

var emp = Empleado{
	Id:           0,
	Nombre:       "Jose",
	Apellido:     "Lorenzo",
	Edad:         19,
	SalarioHora:  2000,
	TipoSalario:  "Hora",
	Complementos: []string{},
	Trabajo: Trabajo{
		Tipo:   "Empleado",
		Sector: "Bailarin de tubo",
		Jefe:   &emp1,
	},
	DiasTrabajo: []DiaTrabajo{{
		Fecha:    time.Time{},
		Duracion: 8 * time.Hour,
	}},
}

func TestCalcularSalario(t *testing.T) {
	salario := emp.CalcularSalario()
	if salario != 16000 {
		t.Fatalf("Expected 0 got %f", salario)
	}
}

func TestObtenerJefe(t *testing.T) {
	gfe := emp.ObtenerJefe()
	if gfe != &emp1 {
		t.Fatalf("Expected a got %s", emp1.Nombre)
	}
}

func TestString(t *testing.T) {
	empleadoString := emp.String()
	if empleadoString != "ID:0,Nombre:Jose,Apellido:Lorenzo,Edad:19" {
		t.Fatalf("Obtuve %s", empleadoString)
	}
}
