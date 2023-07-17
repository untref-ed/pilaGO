package pila

// Pruebas unitarias
import (
	"fmt"
	"testing"
)

func TestPila(t *testing.T) {
	p := NuevaPila[int]()

	p.Push(1)
	p.Push(2)
	p.Push(3)
	for !p.EstaVacia() {
		dato, _ := p.Pop()
		fmt.Println(dato)
	}
	// Output:
	// 3
	// 2
	// 1
}

func TestPilaVacia(t *testing.T) {
	p := NuevaPila[int]()
	_, err := p.Pop()
	fmt.Println(err)
	// Output:
	// Pila vacía
}
