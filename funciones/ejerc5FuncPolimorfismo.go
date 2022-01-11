//interfaces
//crear un tipo cuadrado, circulo y adjuntarle un metodo área
package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}
type cuadrado struct {
	longitud float64
}

//adjuntarle un metodo
func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}
func (c cuadrado) area() float64 {
	return c.longitud * c.longitud
}

type forma interface {
	area() float64
}

func info(f forma) {
	fmt.Println(f.area())
}
func main() {
	circ := circulo{
		radio: 12.345,
	}
	cua := cuadrado{
		longitud: 15,
	}
	info(circ)
	info(cua)
}
