package producto

import "fmt"

type Producto struct {
	Nombre    string
	Cantidad  int
	Categoria string
}

type Insumo struct {
	Nombre   string
	Cantidad int
}

func (p *Producto) IncrementarCantidad(c int) {
	p.Cantidad += c
}

func (p *Producto) DisminuirCantidad(c int) {
	p.Cantidad -= c
}

func (p Producto) String() string {
	return fmt.Sprintf("Nombre:%s, Cantidad:%d, Categoria:%s", p.Nombre, p.Cantidad, p.Categoria)
}
