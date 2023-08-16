package pila

// Pruebas unitarias
import (
	"testing"
)

func TestPila(t *testing.T) {
	p := NuevaPila[int]()

	p.Push(1)
	p.Push(2)
	p.Push(3)

	dato, _ := p.Pop()

	if dato != 3 {
		t.Errorf("Se esperaba 3 y se obtuvo %d", dato)
	}
}

func TestPilaVacia(t *testing.T) {
	p := NuevaPila[int]()

	_, err := p.Pop()

	if err == nil {
		t.Errorf("Se esperaba un error y no se obtuvo")
	}
}
