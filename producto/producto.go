package producto

import "fmt"

type Producto struct {
	NSerie    int
	Nombre    string
	Cantidad  int
	Categoria string
}

type Insumo struct {
	NSerie   int
	Nombre   string
	Cantidad int
}

func (i *Insumo) IncrementarCantidadInsumo(c int) {
	i.Cantidad += c
}

func (i *Insumo) DisminuirCantidadInsumo(c int) {
	i.Cantidad -= c
}

func (p *Producto) IncrementarCantidadProducto(c int) {
	p.Cantidad += c
}

func (p *Producto) DisminuirCantidadProducto(c int) {
	p.Cantidad -= c
}

func (p Producto) String() string {
	return fmt.Sprintf("Nombre:%s, Cantidad:%d, Categoria:%s", p.Nombre, p.Cantidad, p.Categoria)
}
