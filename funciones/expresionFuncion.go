//expresion Funcion
//funciones son ciudadanos de primer orden
//las funciones se pueden asignar a variables
package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("Mi primera expresion funcion")
	}
	f()
	g := func(x int) {
		fmt.Println("El año en que se descubrio america fue: ", x)
	}
	g(1492)
}
