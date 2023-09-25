package producto

import "fmt"

type Producto struct {
	Nombre   string
	Cantidad int
}

func (p Producto) IncrementarCantidad(c int) {
	p.Cantidad = p.Cantidad + c
}

func (p Producto) DisminuirCantidad(c int) {
	p.Cantidad = p.Cantidad + c
}

func (p Producto) String() string {
	return fmt.Sprintf("Nombre de producto:%s, Cantidad:%d", p.Nombre, p.Cantidad)
}
