package producto

import "fmt"

type Producto struct {
	Nombre    string
	Cantidad  int
	Categoria string
}

func (p Producto) IncrementarCantidad(c int) {
	p.Cantidad += c
}

func (p Producto) DisminuirCantidad(c int) {
	p.Cantidad += c
}

func (p Producto) String() string {
	return fmt.Sprintf("Nombre de producto:%s, Cantidad:%d, Categoria:%s\n", p.Nombre, p.Cantidad, p.Categoria)
}
