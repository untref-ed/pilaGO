package pila

import (
	"fmt"
)

// Pila genérica, soporta cualquier tipo de dato
type Pila[T any] struct {
	datos []T
}

// Crea una nueva pila vacía. O(1)
func NuevaPila[T any]() *Pila[T] {
	p := new(Pila[T])
	p.datos = make([]T, 0)
	return p
}

// Apila un dato. O(1)
func (p *Pila[T]) Push(dato T) {
	p.datos = append(p.datos, dato)
}

// Desapila un dato. Si la pila está vacía devuelve un error. O(1)
func (p *Pila[T]) Pop() (T, error) {
	var nulo T
	if len(p.datos) == 0 {
		return nulo, fmt.Errorf("Pila vacía")
	}
	dato := p.datos[len(p.datos)-1]
	p.datos = p.datos[:len(p.datos)-1]
	return dato, nil
}

// Devuelve el dato del tope de la pila. Si la pila está vacía devuelve un error. O(1)
func (p *Pila[T]) Top() (T, error) {
	var nulo T
	if len(p.datos) == 0 {
		return nulo, fmt.Errorf("Pila vacía")
	}
	return p.datos[len(p.datos)-1], nil
}

// Devuelve true si la pila está vacía o false en caso contrario. O(1)
func (p *Pila[T]) EstaVacia() bool {
	return len(p.datos) == 0
}
