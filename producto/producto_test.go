package producto

import (
	"testing"
)

func TestIncrementarCantidad(t *testing.T) {
	producto1 := Producto{
		Nombre:    "Blue Label",
		Cantidad:  12,
		Categoria: "Whisky",
	}

	producto1.IncrementarCantidad(1)
	if !(producto1.Cantidad > 12) {
		t.Fatalf("Expected 12+1 la mama como ninguno got %d", producto1.Cantidad)
	}

}

func TestDisminuirCantidad(t *testing.T) {
	producto1 := Producto{
		Nombre:    "Blue Label",
		Cantidad:  12,
		Categoria: "Whisky",
	}

	producto1.DisminuirCantidad(1)
	if !(producto1.Cantidad < 12) {
		t.Fatalf("Expected 12-1 la mama como ninguno got %d", producto1.Cantidad)
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
