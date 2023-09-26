package producto

import (
	"testing"
)

func TestIncrementarCantidadProducto(t *testing.T) {
	producto1 := Producto{
		Nombre:    "Blue Label",
		Cantidad:  12,
		Categoria: "Whisky",
	}

	producto1.IncrementarCantidadProducto(1)
	if !(producto1.Cantidad > 12) {
		t.Fatalf("Expected 12+1 la mama como ninguno got %d", producto1.Cantidad)
	}

}

func TestDisminuirCantidadProducto(t *testing.T) {
	producto1 := Producto{
		Nombre:    "Blue Label",
		Cantidad:  12,
		Categoria: "Whisky",
	}

	producto1.DisminuirCantidadProducto(1)
	if !(producto1.Cantidad < 12) {
		t.Fatalf("Expected 12-1 la mama como ninguno got %d", producto1.Cantidad)
	}

}

func TestIncrementarCantidadInsumo(t *testing.T) {
	insumo1 := Insumo{
		Nombre:   "Botellas de vidrio",
		Cantidad: 150,
	}

	insumo1.IncrementarCantidadInsumo(1)
	if !(insumo1.Cantidad > 12) {
		t.Fatalf("Expected 12+1 la mama como ninguno got %d", insumo1.Cantidad)
	}

}

func TestDisminuirCantidadInsumo(t *testing.T) {
	insumo1 := Insumo{
		Nombre:   "Botellas de vidrio",
		Cantidad: 150,
	}

	insumo1.DisminuirCantidadInsumo(1)
	if !(insumo1.Cantidad < 150) {
		t.Fatalf("Expected 149 got %d", insumo1.Cantidad)
	}

}

func TestString(t *testing.T) {
	producto1 := Producto{
		Nombre:    "Blue Label",
		Cantidad:  12,
		Categoria: "Whisky",
	}
	productoString := producto1.String()

	if productoString != "Nombre:Blue Label, Cantidad:12, Categoria:Whisky" {
		t.Fatalf("Obtuve %s", productoString)
	}
}
