//Type struct
package main

import "fmt"

type persona struct {
	nombre     string
	apellido   string
	saboresFav []string
}

func main() {
	p1 := persona{
		nombre:   "Charles",
		apellido: "McLean",
		saboresFav: []string{
			"fresa",
			"Chocolate",
			"Ron con pasas",
		},
	}
	p2 := persona{
		nombre:   "Charlize",
		apellido: "McLean McDonald",
		saboresFav: []string{
			"galleta",
			"vainilla",
			"chocolate",
		},
	}
	fmt.Println(p1.nombre)
	fmt.Println(p1.apellido)
	for i, valor := range p2.saboresFav {
		fmt.Println("\t", i, valor)
	}
}
