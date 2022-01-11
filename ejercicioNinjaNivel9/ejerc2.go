//MethodSet
//Si tenemmos receptores metodos de tipo puntero
//Receptores    Valores
//(t T) 		T y *T
//(t *T)		*T

package main

import "fmt"

type Persona struct {
	nombre string
}

func (p *Persona) hablar() {
	fmt.Println("Hola soy", p.nombre)
}

type humano interface {
	hablar()
}

func diAlgo(h humano) {
	h.hablar()
}
func main() {
	p1 := Persona{
		nombre: "Charles",
	}

	diAlgo(&p1) //implementa la interfaz
	//p1.hablar()
}
