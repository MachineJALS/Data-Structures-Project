package main

import (
	"cmp"
	"fmt"
	"time"

	"mymodule2/empleado"
	"mymodule2/producto"
	"slices"
)

var (
	listaEmpleados = []empleado.Empleado{}
	listaProductos = []producto.Producto{}
	idEmpleado     int
)

func main() {
	datosQuemados()
	for {
		menuPrincipal()
	}
}

func menuPrincipal() {
	fmt.Println("=== Menú Principal ===")
	fmt.Println("1) Menú empleados")
	fmt.Println("2) Menú productos")
	fmt.Println("3) Menú insumos")
	fmt.Println("4) Salir del programa")
	fmt.Print("Digite una de las opciones: ")

	var option int
	fmt.Scan(&option)

	switch option {
	case 1:
		fmt.Println("Entrando al menú de empleados...")
		empleadosMenu()
	case 2:
		fmt.Println("Entrando al menú de productos...")
		productosMenu()
	case 3:
		fmt.Println("Entrando al menú de insumos...")
	case 4:
		fmt.Println("Saliendo del programa")
		fmt.Println("¡Hasta luego!")
		// Salir del programa
		panic("Programa terminado")
	default:
		fmt.Println("Opción no válida.")
	}
}

func datosQuemados() {
	empleado1 := empleado.Empleado{
		Id:           1,
		Nombre:       "Oscar",
		Apellido:     "Viquez",
		Edad:         27,
		SalarioHora:  30000,
		TipoSalario:  "Hora",
		Complementos: []string{},
		Trabajo: empleado.Trabajo{
			Tipo:   "Jefe",
			Sector: "Informatico",
		},
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
		Nombre:    "Bud light",
		Cantidad:  130,
		Categoria: "Cerveza",
	}
	producto2 := producto.Producto{
		Nombre:    "Johnni Walker",
		Cantidad:  50,
		Categoria: "Whisky",
	}

	producto3 := producto.Producto{
		Nombre:    "Bud light",
		Cantidad:  132,
		Categoria: "Cerveza",
	}

	listaProductos = append(listaProductos, producto1, producto3, producto2)
	fmt.Printf("Productos:\n%v\n", listaProductos)

	listaEmpleados = append(listaEmpleados, empleado1, empleado2)
	fmt.Printf("Empleados:\n%v\n", listaEmpleados)

	// Ordenar empleados por ID y apellido
	slices.SortFunc(listaEmpleados, func(a, b empleado.Empleado) int { return a.Id - b.Id })
	fmt.Printf("Empleados ordenados por ID:\n%v\n", listaEmpleados)

	slices.SortFunc(listaEmpleados, func(a, b empleado.Empleado) int { return cmp.Compare(a.Apellido, b.Apellido) })
	fmt.Printf("Empleados ordenados por apellido:\n%v\n", listaEmpleados)

	// Encontrar y eliminar un empleado por apellido
	indice := slices.IndexFunc(listaEmpleados, func(e empleado.Empleado) bool {
		return e.Apellido == "Lorenzo"
	})
	listaEmpleados = slices.Delete(listaEmpleados, indice, (indice + 1))
	fmt.Printf("Empleados después de eliminar a Lorenzo:\n%v\n", listaEmpleados)
}

func formatProductos(productos []producto.Producto) string {
	var result string

	for _, p := range productos {
		productoStr := fmt.Sprintf("Nombre: %s, Cantidad: %d, Categoria: %s\n", p.Nombre, p.Cantidad, p.Categoria)
		result += productoStr
	}

	return result
}

func mostrarProductos() {
	fmt.Println("=== Lista de Productos ===")
	productosStr := formatProductos(listaProductos)
	fmt.Print(productosStr)
}

func productosMenu() {
	for {
		fmt.Println("=== Menú de Productos ===")
		fmt.Println("1) Mostrar productos")
		fmt.Println("2) Agregar producto")
		fmt.Println("3) Salir al menú principal")
		fmt.Print("Digite una de las opciones: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			mostrarProductos()
		case 2:
			// Agregar producto (debes implementarlo)
		case 3:
			fmt.Println("Saliendo al menú principal...")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

func empleadosMenu() {
	for {
		fmt.Println("=== Menú de Empleados ===")
		fmt.Println("1) Añadir empleado")
		fmt.Println("2) Eliminar empleado")
		fmt.Println("3) Modificar empleado")
		fmt.Println("4) Salir al menú principal")
		fmt.Print("Digite una de las opciones: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			añadirEmpleado()
		case 2:
			seleccionarEmpleadoEliminar()
		case 3:
			seleccionarEmpleadoModificar()
		case 4:
			fmt.Println("Saliendo al menú principal...")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

func seleccionarEmpleadoModificar() {
	fmt.Println("Lista de empleados:")
	for i, e := range listaEmpleados {
		fmt.Printf("%d) %s %s\n", i+1, e.Nombre, e.Apellido)
	}

	if len(listaEmpleados) == 0 {
		fmt.Println("No hay empleados para modificar.")
		return
	}

	fmt.Print("Elija el número del empleado que desea modificar: ")
	var opcion int
	fmt.Scan(&opcion)

	if opcion >= 1 && opcion <= len(listaEmpleados) {
		index := opcion - 1
		empleadoSeleccionado := &listaEmpleados[index]

		fmt.Println("Empleado seleccionado:")
		fmt.Println(*empleadoSeleccionado)

		fmt.Print("Ingrese el nuevo apellido para el empleado: ")
		var nuevoApellido string
		fmt.Scan(&nuevoApellido)

		// Modificar el apellido del empleado seleccionado
		empleadoSeleccionado.Apellido = nuevoApellido

		fmt.Println("Apellido del empleado modificado con éxito.")
	} else {
		fmt.Println("Opción no válida.")
	}
}

func seleccionarEmpleadoEliminar() {
	fmt.Println("Lista de empleados:")
	for i, e := range listaEmpleados {
		fmt.Printf("%d) %s %s\n", i+1, e.Nombre, e.Apellido)
	}

	if len(listaEmpleados) == 0 {
		fmt.Println("No hay empleados para eliminar.")
		return
	}

	fmt.Print("Elija el número del empleado que desea eliminar: ")
	var opcion int
	fmt.Scan(&opcion)

	if opcion >= 1 && opcion <= len(listaEmpleados) {
		index := opcion - 1
		empleadoSeleccionado := &listaEmpleados[index]

		fmt.Println("Empleado seleccionado:")
		fmt.Println(*empleadoSeleccionado)

		fmt.Print("¿Está seguro de eliminar este empleado? (S/N): ")
		var confirmacion string
		fmt.Scan(&confirmacion)

		if confirmacion == "S" || confirmacion == "s" {
			// Eliminar el empleado seleccionado
			listaEmpleados = append(listaEmpleados[:index], listaEmpleados[index+1:]...)
			fmt.Println("Empleado eliminado con éxito.")
		} else {
			fmt.Println("Operación de eliminación cancelada.")
		}
	} else {
		fmt.Println("Opción no válida.")
	}
}

func añadirEmpleado() {
	var nombre string
	var apellido string
	var edad int
	var tipoSalario string
	var salarioHora float64
	idEmpleado++

	fmt.Print("Ingrese el nombre del empleado: ")
	fmt.Scan(&nombre)

	fmt.Print("Ingrese el apellido del empleado: ")
	fmt.Scan(&apellido)

	fmt.Print("Ingrese la edad del empleado: ")
	fmt.Scan(&edad)

	fmt.Print("Ingrese el salario del empleado: ")
	fmt.Scan(&salarioHora)

	fmt.Print("1) Hora\n2) Mensual\nIngrese el tipo de salario: ")
	var option int
	fmt.Scan(&option)
	if option == 1 {
		tipoSalario = "Hora"
	} else {
		tipoSalario = "Mensual"
	}

	nuevoEmpleado := empleado.Empleado{
		Id:          idEmpleado,
		Nombre:      nombre,
		Apellido:    apellido,
		Edad:        edad,
		SalarioHora: salarioHora,
		TipoSalario: tipoSalario,
		DiasTrabajo: []empleado.DiaTrabajo{},
	}
	listaEmpleados = append(listaEmpleados, nuevoEmpleado)
	fmt.Println("Empleado añadido con éxito!")
}
