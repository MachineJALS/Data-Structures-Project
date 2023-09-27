package main

import (
	"bufio"
	"fmt"
	"mymodule2/empleado"
	"mymodule2/producto"
	"os"
	"sort"
	"time"
)

var (
	listaEmpleados = []empleado.Empleado{}
	listaProductos = []producto.Producto{}
	listaInsumos   = []producto.Insumo{}
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

	insumo1 := producto.Insumo{
		Nombre:   "Botellas",
		Cantidad: 7,
	}

	insumo2 := producto.Insumo{
		Nombre:   "Tapas",
		Cantidad: 9,
	}

	insumo3 := producto.Insumo{
		Nombre:   "Licor",
		Cantidad: 15,
	}

	listaInsumos = append(listaInsumos, insumo1, insumo2, insumo3)

	listaProductos = append(listaProductos, producto1, producto3, producto2)

	listaEmpleados = append(listaEmpleados, empleado1, empleado2)
}

func Supervisor() {

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
		fmt.Println("4) Imprimir empleados por ID")
		fmt.Println("5) Imprimir empleados por Apellido")
		fmt.Println("6) Salir al menú principal")
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
			fmt.Println("Imprimiendo empleados por ID...")
			imprimirEmpleadosOrdenadosPorID()
		case 5:
			fmt.Println("Imprimiendo empleados por Apellido...")
			imprimirEmpleadosOrdenadosPorApellido()
		case 6:
			fmt.Println("Saliendo al menú principal...")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

func imprimirEmpleadosOrdenadosPorID() {
	// Copia la lista de empleados
	empleadosCopia := make([]empleado.Empleado, len(listaEmpleados))
	copy(empleadosCopia, listaEmpleados)

	// Ordena la copia por ID
	sort.Slice(empleadosCopia, func(i, j int) bool {
		return empleadosCopia[i].Id < empleadosCopia[j].Id
	})

	// Imprime la lista ordenada por ID
	fmt.Println("=== Empleados Ordenados por ID ===")
	for i, e := range empleadosCopia {
		fmt.Printf("%d) %s %s (ID: %d)\n", i+1, e.Nombre, e.Apellido, e.Id)
	}
}

func imprimirEmpleadosOrdenadosPorApellido() {
	// Copia la lista de empleados
	empleadosCopia := make([]empleado.Empleado, len(listaEmpleados))
	copy(empleadosCopia, listaEmpleados)

	// Ordena la copia por Apellido
	sort.Slice(empleadosCopia, func(i, j int) bool {
		return empleadosCopia[i].Apellido < empleadosCopia[j].Apellido
	})

	// Imprime la lista ordenada por Apellido
	fmt.Println("=== Empleados Ordenados por Apellido ===")
	for i, e := range empleadosCopia {
		fmt.Printf("%d) %s %s (Apellido: %s)\n", i+1, e.Nombre, e.Apellido, e.Apellido)
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
		reader := bufio.NewReader(os.Stdin)
		nuevoApellido, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Error al leer la entrada:", err)
			return
		}

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

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese el nombre del empleado: ")
	nombre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Error al leer la entrada:", err)
		return
	}

	fmt.Print("Ingrese el apellido del empleado: ")
	apellido, err = reader.ReadString('\n')
	if err != nil {
		fmt.Print("Error al leer la entrada:", err)
	}

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
