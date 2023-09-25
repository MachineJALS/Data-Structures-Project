package main

import (
	"cmp"
	"fmt"
	"mymodule2/empleado"
	"mymodule2/producto"
	"slices"
	"time"
)

var listaEmpleados = []empleado.Empleado{}
var listaProductos = []producto.Producto{}

func main() {
	empleado1 := empleado.Empleado{
		Id:           1,
		Nombre:       "a",
		Apellido:     "L",
		Edad:         1,
		SalarioHora:  2,
		TipoSalario:  "Hora",
		Complementos: []string{},
		Trabajo:      empleado.Trabajo{},
		DiasTrabajo: []empleado.DiaTrabajo{{
			Fecha:    time.Time{},
			Duracion: 8 * time.Hour,
		}},
	}

	empleado2 := empleado.Empleado{
		Id:           0,
		Nombre:       "Jose",
		Apellido:     "Lorenzo",
		Edad:         19,
		SalarioHora:  2000,
		TipoSalario:  "Hora",
		Complementos: []string{},
		Trabajo: empleado.Trabajo{
			Tipo:   "Empleado",
			Sector: "Bailarin de tubo",
			Jefe:   &empleado1,
		},
		DiasTrabajo: []empleado.DiaTrabajo{{
			Fecha:    time.Time{},
			Duracion: 8 * time.Hour,
		}},
	}

	producto1 := producto.Producto{
		Nombre:   "Bud light",
		Cantidad: 130,
	}

	listaProductos = append(listaProductos, producto1)
	fmt.Printf("%v\n", listaProductos)

	listaEmpleados = append(listaEmpleados, empleado1, empleado2)
	fmt.Printf("%v\n", listaEmpleados)

	slices.SortFunc(listaEmpleados, func(a, b empleado.Empleado) int { return a.Id - b.Id })
	fmt.Printf("%v\n", listaEmpleados)

	slices.SortFunc(listaEmpleados, func(a, b empleado.Empleado) int { return cmp.Compare(a.Apellido, b.Apellido) })
	fmt.Printf("%v\n", listaEmpleados)

	indice := slices.IndexFunc(listaEmpleados, func(e empleado.Empleado) bool {
		return e.Apellido == "Lorenzo"
	})
	listaEmpleados = slices.Delete(listaEmpleados, indice, (indice + 1))
	fmt.Printf("%v\n", listaEmpleados)

	//var cantidad int

	//fmt.Printf("Digite la cantidad de elementos que desea añadir a el producto %v :", p)
	//fmt.Scan(&cantidad)

}
