//Sort personalizado
//Ordenación personalizada
package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	Nombre string
	Edad   int
}
type PorEdad []Persona

func (a PorEdad) Len() int           { return len(a) }
func (a PorEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

func main() {
	p1 := Persona{
		Nombre: "Larisha",
		Edad:   9,
	}
	p2 := Persona{
		Nombre: "Shania",
		Edad:   3,
	}
	p3 := Persona{
		Nombre: "Jalisha",
		Edad:   7,
	}
	p4 := Persona{
		Nombre: "Charlize",
		Edad:   1,
	}
	personas := []Persona{p1, p2, p3, p4}
	fmt.Println(personas)

	//ordenar por edad
	sort.Sort(PorEdad(personas))
	fmt.Println(personas)
}
func (p Persona) String() string {
	return fmt.Sprintf("%s: %d", p.Nombre, p.Edad)
}
