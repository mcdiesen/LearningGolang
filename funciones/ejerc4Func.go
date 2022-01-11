//tipo de dato definido por el usuario struct

package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

//adjuntarle un método al struct persona
func (p persona) presentar() {
	fmt.Println("Hola soy", p.nombre, p.apellido, " y tengo ", p.edad, "años.")
}

func main() {
	//crear un valor de tipo persona
	p1 := persona{
		nombre:   "Charles",
		apellido: "McLean",
		edad:     39,
	}
	//llamar el metodo de tipo persona
	p1.presentar()
}
